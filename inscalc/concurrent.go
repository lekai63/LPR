package inscalc

import (
	"fmt"
	"log"
	"sync"
)

func GenAllIns() {
	wg := sync.WaitGroup{}
	items := 13
	wg.Add(items)
	for i := 1; i <= items; i++ {
		go genIns(int32(i), &wg)
	}
	wg.Wait()
}

// wg 一定要通过指针传递
func genIns(i int32, wg *sync.WaitGroup) {
	// ref https://zhuanlan.zhihu.com/p/222367644
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed", err)
		}
	}()

	model, err := NewModel(i)
	if err != nil {
		fmt.Errorf("新建模型时发生错误:%w", err)
	}
	model.ToBank(true)
	fmt.Printf("银行名称:%s 合同号:%s 的还款计划为:", model.Bc.BankName, model.Bc.BankContractNo.ValueOrZero())
	fmt.Print(model.Brps.Table())

	wg.Done()

}
