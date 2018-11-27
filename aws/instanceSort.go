package main

import (
	"sort"

	"github.com/aws/aws-sdk-go/service/ec2"
)

type instancesSlice []*ec2.Instance

func (ins instancesSlice) Len() int {
	return len(ins)
}

func (ins instancesSlice) Less(i, j int) bool {
	iT := ins[i].Tags
	iV := ""
	for _, tag := range iT {
		if *tag.Key == "Name" {
			iV = *tag.Value
		}
	}
	jT := ins[j].Tags
	jV := ""
	for _, tag := range jT {
		if *tag.Key == "Name" {
			jV = *tag.Value
		}
	}
	return iV < jV
}

func (ins instancesSlice) Swap(i, j int) {
	ins[i], ins[j] = ins[j], ins[i]
}
func processResult(result *ec2.DescribeInstancesOutput) instancesSlice {
	var sortedIns instancesSlice
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			sortedIns = append(sortedIns, instance)
		}
	}
	sortInstance(sortedIns)
	return sortedIns
}

func sortInstance(instances instancesSlice) {
	sort.Sort(instances)
}
