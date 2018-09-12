package main

import (
	"fmt"

	"ood/lab3/beverage"
	"ood/lab3/condiment"
	"ood/lab3/decorator"

	"github.com/pkg/errors"
)

func main() {
	getAndPrintBeverageInfo(dialogWithUser)
	getAndPrintBeverageInfo(makeBeverage1)
	getAndPrintBeverageInfo(makeBeverage2)
	getAndPrintBeverageInfo(makeBeverage3)
	getAndPrintBeverageInfo(makeBeverage4)
}

func getAndPrintBeverageInfo(fn func() beverage.Beverage) {
	bev := fn()
	if bev == nil {
		return
	}

	fmt.Printf("%s, cost: %.2f\n", bev, bev.GetCost())
}

func dialogWithUser() beverage.Beverage {
	bev := getBeverageFromDialogWithUser()
	if bev == nil {
		return nil
	}

	for {
		cond, err := getCondimentFromDialogWithUser()
		if err != nil {
			return nil
		}
		if cond == nil {
			return bev
		}

		bev = decorator.Decorate(bev, cond)
	}
}

func getBeverageFromDialogWithUser() beverage.Beverage {
	fmt.Println("1 - Coffee, 2 - Fruit tea")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		return nil
	}

	switch choice {
	case 1:
		return beverage.NewCoffee()
	case 2:
		return beverage.NewTea(beverage.TeaFruit)
	default:
		return nil
	}
}

func getCondimentFromDialogWithUser() (condiment.Condiment, error) {
	fmt.Println("1 - Lemon, 2 - Cinnamon, 0 - Checkout")

	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		return nil, err
	}

	switch choice {
	case 0:
		return nil, nil
	case 1:
		return condiment.NewLemon(2), nil
	case 2:
		return condiment.NewCinnamon(), nil
	default:
		return nil, errors.Errorf("Unexpected choice `%d`", choice)
	}

}

func makeBeverage1() beverage.Beverage {
	latte := beverage.NewLatte(beverage.LatteNormal)
	withCinnamon := decorator.Decorate(latte, condiment.NewCinnamon())
	withLemon := decorator.Decorate(withCinnamon, condiment.NewLemon(2))
	withIceCubes := decorator.Decorate(withLemon, condiment.NewIceCubes(condiment.DryIce, 2))
	return decorator.Decorate(withIceCubes, condiment.NewChocolateCrumbs(2))
}

func makeBeverage2() beverage.Beverage {
	return decorator.Decorate(
		decorator.Decorate(
			decorator.Decorate(
				decorator.Decorate(
					beverage.NewMilkshake(beverage.MilkshakeMedium),
					condiment.NewCinnamon()),
				condiment.NewLemon(2)),
			condiment.NewIceCubes(condiment.DryIce, 2)),
		condiment.NewChocolateCrumbs(2),
	)
}

func makeBeverage3() beverage.Beverage {
	lemon2 := decorator.MakeDecoratorFunc(condiment.NewLemon(2))
	iceCubes3 := decorator.MakeDecoratorFunc(condiment.NewIceCubes(condiment.WaterIce, 3))
	coconutFlakes := decorator.MakeDecoratorFunc(condiment.NewCoconutFlakes(42))
	tea := beverage.NewTea(beverage.TeaGreen)

	return coconutFlakes(iceCubes3(lemon2(tea)))
}

func makeBeverage4() beverage.Beverage {
	return decorator.NewBuilder(beverage.NewCapuccino(beverage.CapuccinoNormal)).
		WithCondiment(condiment.NewCinnamon()).
		WithCondiment(condiment.NewLemon(2)).
		WithCondiment(condiment.NewSyrup(condiment.MapleSyrup)).
		WithCondiment(condiment.NewChocolateCrumbs(2)).
		Build()
}
