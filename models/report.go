package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type NodeValue struct {
	Volume float32 `json:"Volume"` //数量
	N      int     `json:"N"`      //单数
}
type Node struct {
	Key    string
	Volume float32 `json:"Volume"` //数量
	N      int     `json:"N"`      //单数
}

type PlayerReport struct {
	Expense Node `json:"Expense"`
	Prize   Node `json:"Prize"`
	//Deposit     Node
	//Withdraw    Node
	//Return      float32         //返点
	//Other       float32         //活动
	LotteryStat []Node `json:"LotteryStat"` //彩种名：node
	PlayStat    []Node `json:"PlayStat"`    //玩法名：node
	Hour        []Node `json:"Hour"`
}

//每天的Daily.txt文件信息，也作为一个月的统计信息的回吐数据结构
type DailyReport struct {
	PlayerReport
	NumPlayer int
}

type Report struct {
	User        map[string]PlayerReport `json:User`
	BalanceStat map[string]Node
	Daily       DailyReport
	Date        string
}

//获取reportId的结构
type ReportObject struct {
	ReportId string
}

var (
	MonthReports map[string]*DailyReport //多个月的统计信息的回吐数据map
)

const (
	reserve   = iota
	January   = iota
	February  = iota
	March     = iota
	April     = iota
	May       = iota
	June      = iota
	July      = iota
	August    = iota
	September = iota
	October   = iota
	November  = iota
	December  = iota
)

func init() {
	MonthReports = make(map[string]*DailyReport)
}

//根据reportId即MMMMYY格式的日期字符串进行一次统计
func AddOne(report ReportObject) string {
	if report.ReportId == "" {
		return ""
	}
	//获取report.ReportId标识的那个月的Daily.txt文件内容，存储到DailyDatas中
	DailyDatas := ReadMonth(report.ReportId)

	//根据DailyDatas中存储的内容进行数据统计
	StatisticsMonth(report.ReportId, DailyDatas)

	return report.ReportId
}

//根据ReportId进行一次数据GET操作。ReportId标识一个月
func GetOne(ReportId string) (dailyreport *DailyReport, err error) {
	if v, ok := MonthReports[ReportId]; ok {
		return v, nil
	}
	return nil, errors.New("ReportId Not Exist")
}

//GET所有统计数据
func GetAll() map[string]*DailyReport {
	return MonthReports
}

//根据ReportId删除一个月的统计数据
func Delete(ReportId string) {
	delete(MonthReports, ReportId)
}

//判断当前年是否为闰年
func IsLeapYear(nYear int64) bool {
	if nYear%100 == 0 {
		if nYear%400 == 0 {
			return true
		} else {
			return false
		}
	} else {
		if nYear%4 == 0 {
			return true
		}
	}
	return false
}

//获取每天的Daily.txt的Data,并格式化到结构体中
func ReadDaily(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		return ""
		//panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))

	return string(fd)
}

//获取report.ReportId标识的那个月的Daily.txt文件内容，存储到DailyDatas中
func ReadMonth(path string) map[string]*DailyReport {

	DailyDatas := make(map[string]*DailyReport)
	if len(path) != 6 {
		fmt.Println("month does not exist.")
		return nil
	}

	spath := []byte(path)
	year := spath[:4]
	month := spath[4:]
	nyear, _ := strconv.ParseInt(string(year), 10, 64)   //获取年的数值
	nmonth, _ := strconv.ParseInt(string(month), 10, 64) //获取月的数值

	fmt.Println(path)

	if nmonth >= January && nmonth <= December {
		switch nmonth {
		case January, March, May, July, August, October, December: //大月31天
			{
				//fmt.Println("this month have 31 days.")
				for i := 1; i <= 31; i++ {
					daypath := fmt.Sprintf("%4s/%2s/%02d", year, month, i)

					strDaily := ReadDaily(fmt.Sprintf("./%s/Daily.txt", daypath))
					if strDaily != "" {
						var s DailyReport
						json.Unmarshal([]byte(strDaily), &s)
						DailyDatas[daypath] = &s
					}
				}
			}
		case February:
			{
				if IsLeapYear(nyear) { //闰年2月29天
					//fmt.Println("this month have 29 days.")
					for i := 1; i <= 29; i++ {
						daypath := fmt.Sprintf("%4s/%2s/%02d", year, month, i)

						strDaily := ReadDaily(fmt.Sprintf("./%s/Daily.txt", daypath))
						if strDaily != "" {
							var s DailyReport
							json.Unmarshal([]byte(strDaily), &s)
							DailyDatas[daypath] = &s
						}
					}
				} else { //平年2月28天
					//fmt.Println("this month have 28 days.")
					for i := 1; i <= 28; i++ {
						daypath := fmt.Sprintf("%4s/%2s/%02d", year, month, i)

						strDaily := ReadDaily(fmt.Sprintf("./%s/Daily.txt", daypath))
						if strDaily != "" {
							var s DailyReport
							json.Unmarshal([]byte(strDaily), &s)
							DailyDatas[daypath] = &s
						}
					}
				}
			}
		case April, June, September, November: //小月30天
			{
				//fmt.Println("this month have 30 days.")
				for i := 1; i <= 30; i++ {
					daypath := fmt.Sprintf("%4s/%2s/%02d", year, month, i)

					strDaily := ReadDaily(fmt.Sprintf("./%s/Daily.txt", daypath))
					if strDaily != "" {
						var s DailyReport
						json.Unmarshal([]byte(strDaily), &s)
						DailyDatas[daypath] = &s
					}
				}
			}
		default:
			{
				fmt.Println("should not appear this.")
				return nil
			}
		}
	}

	return DailyDatas
}

//根据DailyDatas中存储的内容进行数据统计
func StatisticsMonth(reportId string, DailyDatas map[string]*DailyReport) {
	var s DailyReport

	ReportLotteryStatMap := make(map[string]*NodeValue) //彩种命map
	ReportPlayStatMap := make(map[string]*NodeValue)    //玩法名map
	ReportHourMap := make(map[string]*NodeValue)        //小时map

	for _, v := range DailyDatas { //迭代每天的Daily.txt数据
		for _, v4 := range v.LotteryStat { //迭代每天的Daily.txt中的LotteryStat数据
			//查询LotteryStat map中是否存在当前获取到的LotteryStat名称
			_, ok := ReportLotteryStatMap[v4.Key]
			if ok { //如果存在累加量和单数
				ReportLotteryStatMap[v4.Key].Volume += v4.Volume
				ReportLotteryStatMap[v4.Key].N += v4.N
			} else { //如果不存在则将该LotteryStat Node加入map
				var nValue NodeValue
				nValue.Volume = v4.Volume
				nValue.N = v4.N
				ReportLotteryStatMap[v4.Key] = &nValue
			}

		}

		for _, v5 := range v.PlayStat { //迭代每天的Daily.txt中的PlayStat数据
			//查询PlayStat map中是否存在当前获取到的PlayStat名称
			_, ok := ReportPlayStatMap[v5.Key]
			if ok { //如果存在累加量和单数
				ReportPlayStatMap[v5.Key].Volume += v5.Volume
				ReportPlayStatMap[v5.Key].N += v5.N
			} else { //如果不存在则将该PlayStat Node加入map
				var nValue NodeValue
				nValue.Volume = v5.Volume
				nValue.N = v5.N
				ReportPlayStatMap[v5.Key] = &nValue
			}
		}
		for _, v6 := range v.Hour { //迭代每天的Daily.txt中的Hour数据
			//查询Hour map中是否存在当前获取到的Hour名称
			_, ok := ReportHourMap[v6.Key]
			if ok { //如果存在累加量和单数
				ReportHourMap[v6.Key].Volume += v6.Volume
				ReportHourMap[v6.Key].N += v6.N
			} else { //如果不存在则将该Hour Node加入map
				var nValue NodeValue
				nValue.Volume = v6.Volume
				nValue.N = v6.N
				ReportHourMap[v6.Key] = &nValue
			}
		}
		//累加NumPlayer
		s.NumPlayer += v.NumPlayer
	}
	//从LotteryStat map中获取LotteryStat Node数据加入DailyReport结构的LotteryStat slice中
	for k7, v7 := range ReportLotteryStatMap {
		var n Node
		n.Key = k7
		n.Volume = v7.Volume
		n.N = v7.N
		s.PlayerReport.LotteryStat = append(s.PlayerReport.LotteryStat, n)
	}
	//从PlayStat map中获取PlayStat Node数据加入DailyReport结构的PlayStat slice中
	for k8, v8 := range ReportPlayStatMap {
		var n Node
		n.Key = k8
		n.Volume = v8.Volume
		n.N = v8.N
		s.PlayerReport.PlayStat = append(s.PlayerReport.PlayStat, n)
	}
	//从Hour map中获取Hour Node数据加入DailyReport结构的Hour slice中
	for k9, v9 := range ReportHourMap {
		var n Node
		n.Key = k9
		n.Volume = v9.Volume
		n.N = v9.N
		s.PlayerReport.Hour = append(s.PlayerReport.Hour, n)
	}
	//以reportId为key将当月的report数据加入map
	MonthReports[reportId] = &s

	return
}
