package core

import (
	"strconv"
	"queue/model"

)

//将job id 放入篮子中
//实际上bucket里面的内容是有序切不重复的
func pushBucket(key string, delayTime int, jobId int) error {
	_, err := exec("ZADD", key, delayTime, jobId)
	return err
}

//从bucket中获取数据()
func getDataFromBucket(key string) (* model.BucketItem, error) {
	res, err := exec("ZRANGE", key, 0, 0, "WITHSCORES")
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	var valueBytes []interface{}
	valueBytes = res.([]interface{})
	//change byte to string
	if len(valueBytes) == 0 {
		return nil, nil
	}
	timestampStr := string(valueBytes[1].([]byte))
	jobIdStr := string(valueBytes[0].([]byte))
	//add a bucket
	item := &model.BucketItem{}
	item.Timestamp, _ = strconv.Atoi(timestampStr)
	item.Jobid, _ = strconv.Atoi(jobIdStr)
	return item, nil
}

//从篮子里面删除该数据
func removeFromBucket(bucketName string, id int) error {
	_, err := exec("ZREM", bucketName, id)
	return err
}
