package circumstance
import(
  "math/rand"
)

type Circumstance struct {
  Influence: float
  Description: string
}

GenerateCircumstances() {
  chances := []

  for i := 0; i < 100; i++ {
    rand.Seed(i)
    flip := rand.Intn(100)

    infl := rand.Float64() + 1
    desc : = "Pretty good sales today."

    if flip % 4 = 0 {
      infl = infl * -1
      desc = "Sales where a little soft today."
    }

    chance := Circumstance {
      Influence: infl,
      Description: desc
    }
  }
}
