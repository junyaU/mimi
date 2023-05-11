package recipe

import "errors"

type cookingTime int

const (
	fiveM cookingTime = iota + 1
	//tenM
	//fifteenM
	twentyM
	//thirtyM
	//fortyFiveM
	//oneH
	//twoH
	//threeH
	//fourH
	fiveH
)

func newCookingTime(value int) (cookingTime, error) {
	c := cookingTime(value)
	if err := c.checkValidTime(); err != nil {
		return 0, err
	}

	return c, nil
}

func (c cookingTime) String() string {
	values := [...]string{"", "fiveM", "tenM", "fifteenM", "twentyM", "thirtyM", "fortyFiveM", "oneH", "twoH", "threeH", "fourH", "fiveH"}
	return values[c]
}

func (c cookingTime) checkValidTime() error {
	if fiveM > c || fiveH < c {
		return errors.New("有効な時間が指定されていません")
	}

	return nil
}
