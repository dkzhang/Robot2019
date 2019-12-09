package client

import (
	"Robot2019/cache"
	"Robot2019/dataServer/configuration"
	"Robot2019/myUtil"
	"Robot2019/webServer/inspectionRecord/realtimeRecord"
	"context"
	"fmt"
	"log"
	"time"

	pb "Robot2019/dataServer/thermalImaging/grpc"
	"google.golang.org/grpc"
)

func ThermalImaging(inspectionID, recordID int) (err error) {

	address := configuration.ThermalImaging_ADDRESS
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Printf(" fatal error! did not connect: %v", err)
		return fmt.Errorf("grpc CollectThermalImagingData grpc.Dial error: %v", err)
	}
	log.Printf("grpc.Dial OK!")
	defer conn.Close()

	c := pb.NewThermalImagingServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	log.Printf("context.WithTimeout() OK!")
	defer cancel()

	r, err := c.CollectRenderAnalyze(ctx, &pb.ThermalImagingRequest{Tag: myUtil.FormatTime(time.Now())})
	if err != nil {
		log.Printf(" fatal error! could not reply: %v", err)
		return fmt.Errorf("grpc CollectThermalImagingData Reply error: %v", err)
	}
	log.Printf("reply = %v", r)

	rti := realtimeRecord.RealTimeInfo{
		InspectionID: inspectionID,
		RecordID:     recordID,
		DateTime:     myUtil.FormatTime(time.Now()),
		Level:        r.Level,
		TextContent:  r.AnalysisReport,
		ImageUrl:     configuration.ThermalImageFile_WEB_URL + r.Filepath + r.Filename,
	}

	err = WriteThermalImagingRecord(&rti)
	if err != nil {
		return fmt.Errorf("WriteThermalImagingRecord error: %v", err)
	}

	return nil
}

func WriteThermalImagingRecord(prti *realtimeRecord.RealTimeInfo) (err error) {
	//将结果写入redis
	opts := &cache.RedisOpts{
		Host: cache.RedisHost,
	}
	theRedis := cache.NewRedis(opts)

	redisKey := fmt.Sprintf("Inspection::%d", prti.InspectionID)
	err = theRedis.ListRPush(redisKey, *prti)
	if err != nil {
		return fmt.Errorf("theRedis.ListRPush error: %v", err)
	}

	//InspectionID SET
	//检查当前轮次，只保留最近三次record数组
	//也即：删除 current-2 ~ new-3的record数组
	currentInspectionID := theRedis.Get("CurrentInspectionID")
	if currentInspectionID != nil {
		ciID, ok := currentInspectionID.(int)
		if ok {
			for i := ciID - 2; i <= prti.InspectionID; i++ {
				theRedis.Delete(fmt.Sprintf("Inspection::%d", i))
			}
		}
	}
	err = theRedis.Set("CurrentInspectionID", prti.InspectionID, time.Hour*24)
	if err != nil {
		return fmt.Errorf("theRedis.Set error: %v", err)
	}
	return nil
}
