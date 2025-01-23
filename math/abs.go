package math

func AbsInt(n1 int, n2 int) int {
  distance := n1 - n2
  if (distance < 0) {
    return -distance
  }
  return distance
}
