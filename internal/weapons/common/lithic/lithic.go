package lithic

import (
	"fmt"

	"github.com/genshinsim/gsim/pkg/core"
)

func init() {
	core.RegisterWeaponFunc("lithic spear", weapon)
	core.RegisterWeaponFunc("lithic blade", weapon)
}

//For every character in the party who hails from Liyue, the character who equips this
//weapon gains 6/7/8/9//10% ATK increase and 2/3/4/5/6% CRIT Rate increase.
func weapon(char core.Character, c *core.Core, r int, param map[string]int) {

	stacks := 0

	c.Events.Subscribe(core.OnInitialize, func(args ...interface{}) bool {
		for _, char := range c.Chars {
			if char.Zone() == core.ZoneLiyue {
				stacks++
			}
		}
		return true
	}, fmt.Sprintf("lithic-%v", char.Name()))

	val := make([]float64, core.EndStatType)
	val[core.CR] = (0.02 + float64(r)*0.01) * float64(stacks)
	val[core.ATKP] = (0.06 + float64(r)*0.01) * float64(stacks)

	char.AddMod(core.CharStatMod{
		Key:    "lithic",
		Expiry: -1,
		Amount: func(a core.AttackTag) ([]float64, bool) {
			return val, true
		},
	})
}