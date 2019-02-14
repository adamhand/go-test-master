package main

import "fmt"

/*
收益接口。包含了两个方法：calculate() 计算并返回项目的收入，
而 source() 返回项目名称。
 */
type Income interface {
	calculate() int
	source() string
}

/*
FixedBilling 项目的结构体类型
有两个字段：projectName 表示项目名称，而 biddedAmount 表示组织向该项目投标的金额。
 */
type FixedBilling struct {
	projectName string
	biddedAmount int
}

/*
TimeAndMaterial项目结构体类型
拥有三个字段名：projectName、noOfHours 和 hourlyRate。
 */
type TimeAndMaterial struct {
	projectName string
	noOfHours int
	hourlyRate int
}

type Advertisement struct {
	adName string
	CPC int
	noOfClicks int
}


/*
实现接口
 */
func (fb FixedBilling)calculate() int {
	return fb.biddedAmount
}

func (fb FixedBilling)source() string {
	return fb.projectName
}

func (tm TimeAndMaterial)calculate() int {
	return tm.hourlyRate * tm.noOfHours
}

func (tm TimeAndMaterial)source() string {
	return tm.projectName
}

func (ad Advertisement)calculate() int {
	return ad.noOfClicks * ad.CPC
}

func (ad Advertisement)source() string {
	return ad.adName
}

/*
计算收益
根据 Income 接口的具体类型，程序会调用不同的 calculate() 和 source() 方法
于是，在 calculateNetIncome 函数中就实现了多态。
 */
func calculateIncome(ic []Income)  {
	var totalIncome int = 0
	for _, income := range ic{
		fmt.Printf("income from %s is $%d\n", income.source(), income.calculate())
		totalIncome += income.calculate()
	}
	fmt.Printf("total income is $%d\n", totalIncome)
}

func main()  {
	project1 := FixedBilling{projectName: "Project 1", biddedAmount: 5000}
	project2 := FixedBilling{projectName: "Project 2", biddedAmount: 10000}
	project3 := TimeAndMaterial{projectName: "Project 3", noOfHours: 160, hourlyRate: 25}
	bannerAd := Advertisement{adName: "Banner Ad", CPC: 2, noOfClicks: 500}
	popupAd := Advertisement{adName: "Popup Ad", CPC: 5, noOfClicks: 750}
	incomeStreams := []Income{project1, project2, project3, bannerAd, popupAd}
	calculateIncome(incomeStreams)
}
