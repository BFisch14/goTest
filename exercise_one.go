package go_unit_test_bootcamp

func FindMissingDrone(droneIds []int) int {
    droneMap := make(map[int]int)
    for i := range droneIds {
        droneMap[droneIds[i]]++
    }
 
    for k, v := range droneMap {
        if v == 1 {
          return k
        }
    }
    return -1
}
