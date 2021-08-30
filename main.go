package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation(), client.WithHost("http://192.168.54.134:2375"))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	var stats types.ContainerStats

	for _, container := range containers {
		fmt.Println(container.ID)
		fmt.Println(container.Names)

		stats, err = cli.ContainerStatsOneShot(ctx, container.ID)
		if err != nil {
			panic(err)
		}

		var parsedStats types.Stats

		json.NewDecoder(stats.Body).Decode(&parsedStats)
		fmt.Println(parsedStats.MemoryStats)
	}
	defer stats.Body.Close()
}
//
//type (
//	throttlingData struct {
//		Periods int64 `json:"periods"`
//		ThrottledPeriods int64 `json:"throttled_periods"`
//		ThrottledTime int64 `json:"throttled_time"`
//	}
//
//	cpuUsage struct {
//		TotalUsage int64 `json:"total_usage"`
//		PercpuUsage []int64 `json:"percpu_usage"`
//		UsageInKernelmode int64 `json:"usage_in_kernelmode"`
//		UsageInUsermode int64 `json:"usage_in_usermode"`
//	}
//
//	cpuStats struct {
//		CpuUsage cpuUsage `json:"cpu_usage"`
//		SystemCpuUsage int64 `json:"system_cpu_usage"`
//		OnlineCpus int64 `json:"online_cpus"`
//		types.ThrottlingData
//	}
//
//	statsResult struct {
//		Time string `json:"read"`
//		LastTime string `json:"preread"`
//		Unknow string `json:"pids_stats"`
//		Unknow string `json:"pids_stats"`
//	}
//
//	statsResponse struct {
//		ResponseData struct{
//			Results []myResult `json:"results"`
//		} `json:"responseData"`
//	}
//)