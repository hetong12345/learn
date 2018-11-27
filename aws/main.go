package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const BJ = "cn-north-1"
const NX = "cn-northwest-1"

var (
	region       string
	instanceId   string
	tag          string
	inputName    string
	full         bool
	state        bool
	outputId     bool
	help         bool
	showStateAll bool
	noName       bool
	noIP         bool
	remoteTagMap map[string]string
	localTagMap  map[string]string
)

func flagParse() {
	flag.StringVar(&region, "r", "bj", "select region from bj/nx/beijing/ningxia/BeiJing/...")
	flag.StringVar(&region, "region", "bj", "equal to -r")
	flag.StringVar(&instanceId, "info", "", "input instance ID , output Instance all info.(Exact match)")
	flag.StringVar(&inputName, "name", "", "equal to -n")
	flag.StringVar(&inputName, "n", "", "matching instance by Name, filter Instance by Instance name.(Fuzzy matching)")
	flag.StringVar(&tag, "t", "", "tag ,filter instance by tag's key,--all output all tags (Ignore case)(Fuzzy matching)")
	flag.StringVar(&tag, "tag", "", "equal to -t")
	flag.BoolVar(&full, "a", false, "all , output complete info")
	flag.BoolVar(&outputId, "id", false, "instance Id ,additional output instanceId")
	flag.BoolVar(&state, "s", false, "state ,additional output Instance state")
	flag.BoolVar(&state, "state", false, "equal to -s")
	flag.BoolVar(&help, "h", false, "help , output mhost help page")
	flag.BoolVar(&noName, "noname", false, "don't output instance name")
	flag.BoolVar(&noIP, "noip", false, "don't output instance IP")
	flag.BoolVar(&showStateAll, "showStateAll", false, "output include any state instance.(default just output state='running' instance)")
	flag.Parse()
}

func main() {

	flagParse()

	if help {
		outputHelp()
		os.Exit(0)
	}

	caseRegion()

	getInfo()

}

func caseRegion() {
	region = strings.ToLower(region)
	switch region {
	case "bj":
		region = BJ
	case "beijing":
		region = BJ
	case "nx":
		region = NX
	case "ningxia":
		region = NX
	default:
		fmt.Println("invalid region , auto exit")
		os.Exit(2)
	}
}

func tagHandle() {
	tagSlice := strings.Split(tag, ",")
	for _, value := range tagSlice {
		index := strings.Index(value, "=")
		if index != -1 {
			localTagMap[strings.ToLower(value[0:index])] = strings.ToLower(value[index+1:])
		} else {
			localTagMap[value] = "*"
		}
	}
}

func outputHelp() {
	fmt.Println("USAGE:")
	fmt.Println("   mhost [options] [-r/region bj/nx] [-s/state] [-n/name] [-id] [-a] " +
		"[-t key=value,key2=value2] [-nname] [-nip] [-showStateAll] [-info xxx]")
	fmt.Println("VERSION:")
	fmt.Println("   mhost v1.0")
	fmt.Println("OPTIONS:")
	flag.PrintDefaults()
	fmt.Println("EXAMPLE:")
	fmt.Println("   mhost")
	fmt.Println("   mhost -a")
	fmt.Println("   mhost -r nx")
	fmt.Println("   mhost -r nx -s")
	fmt.Println("   mhost -r nx -s -nname")
	fmt.Println("   mhost -r nx -s -nip")
	fmt.Println("   mhost -r nx -s -id")
	fmt.Println("   mhost -r nx -s -id -t dep")
	fmt.Println("   mhost -r bj -s -id -t dep=live,prom=tr")
	fmt.Println("   mhost -r bj -s -id -t dep=live,prom=tr -n nginx")
	fmt.Println("   mhost -r bj -s -id -name gunicorn")
	fmt.Println("   mhost -r bj -s -id -name gunicorn -showStateAll")
	fmt.Println("   mhost -r bj -info i-0c74f5c4d502b270a")
}

func newSession(region string) (*session.Session, error) { //get this region session
	return session.NewSession(&aws.Config{Region: aws.String(region)})
}

func getInfo() { //get aws resp
	sess, err := newSession(region)
	if err != nil {
		fmt.Println("failed to create session,", err)
		os.Exit(2)
	}
	// Create an EC2 service client.
	svc := ec2.New(sess)
	input := &ec2.DescribeInstancesInput{}

	result, err := svc.DescribeInstances(input)
	if err != nil {
		fmt.Println("failed to get IP,", err)
		os.Exit(2)
	}

	if instanceId != "" {
		formatPrint(getInstanceInfo(result))
		os.Exit(0)
	}

	localTagMap = make(map[string]string)
	tagHandle()

	sortedInstances := processResult(result)

	if inputName != "" {
		instances := getInstanceByName(sortedInstances)
		if instances == nil {
			os.Exit(2)
		}
		for _, instance := range instances {
			formatPrint(fmtReservation(instance))
		}
		os.Exit(0)
	}

	for _, instance := range sortedInstances {
		if showStateAll {
			formatPrint(fmtReservation(instance))
		} else if *instance.State.Name == "running" {
			formatPrint(fmtReservation(instance))
		}
	}

}

func formatPrint(str string) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 40, 0, 0, ' ', 0)
	fmt.Fprint(w, str)
	w.Flush()
}

func getInstanceByName(ins []*ec2.Instance) []*ec2.Instance {
	var instances []*ec2.Instance
	for _, instance := range ins {
		for _, value := range instance.Tags {
			if aws.StringValue(value.Key) == "Name" {
				if strings.Contains(aws.StringValue(value.Value), inputName) {
					if showStateAll {
						instances = append(instances, instance)
					} else if *instance.State.Name == "running" {
						instances = append(instances, instance)
					}
				}
			}

		}
	}
	return instances
}
func getInstanceInfo(result *ec2.DescribeInstancesOutput) string {
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			if *instance.InstanceId == instanceId {
				return fmt.Sprint(instance)
			}
		}
	}
	return ""
}

func appendIP(ins *ec2.Instance) string {
	return fmt.Sprintf("%s", aws.StringValue(ins.NetworkInterfaces[0].PrivateIpAddress))
}

func appendId(ins *ec2.Instance) string {
	return fmt.Sprintf("\t%s", aws.StringValue(ins.InstanceId))
}

func appendName(ins *ec2.Instance) string {
	for _, value := range ins.Tags {
		if aws.StringValue(value.Key) == "Name" {
			if noIP {
				return fmt.Sprintf("%s", aws.StringValue(value.Value))
			} else {
				return fmt.Sprintf("\t%s", aws.StringValue(value.Value))
			}
		}
	}
	return ""
}

func appendTag(ins *ec2.Instance) string {
	ret := ""
	if tag == "--all" || full {
		for _, remoteTag := range ins.Tags {
			ret += fmt.Sprintf("\tTag:%s=%s", aws.StringValue(remoteTag.Key), aws.StringValue(remoteTag.Value))
		}
	} else {
		for localKey, localValue := range localTagMap {
			for _, remoteTag := range ins.Tags {
				if strings.Contains(strings.ToLower(aws.StringValue(remoteTag.Key)), localKey) &&
					(strings.Contains(strings.ToLower(aws.StringValue(remoteTag.Value)), localValue) || localValue == "*") {
					ret += fmt.Sprintf("\tTag:%s=%s", aws.StringValue(remoteTag.Key), aws.StringValue(remoteTag.Value))
					break
				}
			}
		}
	}
	return ret
}

func appendState(ins *ec2.Instance) string {
	return fmt.Sprintf("\t%s", aws.StringValue(ins.State.Name))
}

func fmtReservation(ins *ec2.Instance) string {
	out := ""
	if !noIP {
		out += appendIP(ins)
	}
	if !noName {
		out += appendName(ins)
	}

	if full { //加了-a 则忽略其他输出
		out += appendState(ins)
		out += appendId(ins)
		out += appendTag(ins)
		return out + "\n"
	}

	if tag != "" {
		if tag == "--all" { //需要输出全部tag
			if state {
				out += appendState(ins)
			}

			if outputId {
				out += appendId(ins)
			}
			out += appendTag(ins)
			return out + "\n"
		} else {
			if instanceFilter(ins) { //判断当前instance是否需要输出
				if state {
					out += appendState(ins)
				}
				if outputId {
					out += appendId(ins)
				}
				out += appendTag(ins)
				return out + "\n"
			} else {
				return ""
			}
		}
	}

	if state {
		out += appendState(ins)
	}

	if outputId {
		out += appendId(ins)
	}

	if out != "" {
		out += "\n"
	}

	return out
}

func instanceFilter(ins *ec2.Instance) bool {
	remoteTagMap = make(map[string]string)
	//如果指定了标签，则分割成切片
	for _, tag := range ins.Tags {
		remoteTagMap[strings.ToLower(aws.StringValue(tag.Key))] = strings.ToLower(aws.StringValue(tag.Value))
	}
	allContains := true
	for localKey, localValue := range localTagMap { //双层循环确认tag是否都存在
		contains := false
		for remoteKey, remoteValue := range remoteTagMap {
			if strings.Contains(remoteKey, localKey) {
				if localValue == "*" || strings.Contains(remoteValue, localValue) {
					contains = true
					break
				}
			}
		}
		if !contains {
			allContains = false
		}
	}
	return allContains
}
