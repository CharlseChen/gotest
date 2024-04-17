package benchmark

import (
	"testing"
	"github.com/gogf/gf/util/grand"
)

func Benchmark_ActCarnivalGift(b *testing.B) {
	var hasBuffDouble bool
	hasBuffDouble = false
	awardNum := 0
	gid := 11017286
	awardCid := 6284
	//检测是否前30次没有获得戒指，如果没有获得，此次必中

	//gid := 11017284
	//awardCid := 6283
	var ringCount uint32 = 1

	for j := 0; j < b.N; j++ {
		var count uint32 = 38
		//前五次的返利概率
		var weight, totalWeight uint32
		//便宜戒指
		if gid == 11017284 {
			if count <= 5 {
				weight = 40
				totalWeight = 240
				if hasBuffDouble {
					weight = 50
					totalWeight = 250
				}
			} else if 5 < count {
				weight = 20
				totalWeight = 220
				if hasBuffDouble {
					weight = 30
					totalWeight = 230
				}
				//if count >= 30 && ringCount == 0 && awardNum == 0 {
				//	weight = 100
				//	totalWeight = 100
				//}
			}
			//贵戒指
		} else if gid == 11017286 {
			if count <= 5 {
				weight = 15
				totalWeight = 215
				if hasBuffDouble {
					weight = 23
					totalWeight = 223
				}
			} else if 5 < count {
				weight = 10
				totalWeight = 210
				if hasBuffDouble {
					weight = 20
					totalWeight = 220
				}
				if count >= 38 && ringCount == 0 && awardNum == 0 {
					weight = 100
					totalWeight = 100
				}
			}
		}
		if int(weight) >= grand.N(1, int(totalWeight)) {
			awardNum++
		}
	}
	//发一天戒指奖励
	b.Logf("轮次：%d, 有buff否:%v, 奖励的cid:%d, 奖励的数量:%d", b.N, hasBuffDouble, awardCid, awardNum)
}
