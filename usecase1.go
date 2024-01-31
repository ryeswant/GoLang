package main

import "fmt"

func addNewFamily(productFamily map[string][]string) {
	var familyKey string

	fmt.Print("Enter a new Product family key: ")
	fmt.Scanln(&familyKey)

	if _, ok := productFamily[familyKey]; ok {
		fmt.Printf("%s already exists in the map\n", familyKey)
	} else {
		productFamily[familyKey] = []string{}
		fmt.Printf("%s added to the map\n", familyKey)
	}
}

func addNodeToFamily(productFamily map[string][]string) {
	var familyKey, nodeName string

	fmt.Print("Enter the family key: ")
	fmt.Scanln(&familyKey)

	if nodes, ok := productFamily[familyKey]; ok {
		fmt.Print("Enter the node name to add to family ", familyKey, ": ")
		fmt.Scanln(&nodeName)

		if contains(nodes, nodeName) {
			fmt.Printf("Node %s already exists in family %s\n", nodeName, familyKey)
		} else {
			productFamily[familyKey] = append(nodes, nodeName)
			fmt.Printf("Node %s added to family %s\n", nodeName, familyKey)
		}
	} else {
		fmt.Printf("Family %s does not exist in the map\n", familyKey)
	}
}

func deleteNodeFromFamily(productFamily map[string][]string) {
	var familyKey, nodeName string

	fmt.Print("Enter the family key: ")
	fmt.Scanln(&familyKey)

	if nodes, ok := productFamily[familyKey]; ok {
		fmt.Print("Enter the node name to delete from family ", familyKey, ": ")
		fmt.Scanln(&nodeName)

		if !contains(nodes, nodeName) {
			fmt.Printf("Node %s not found in family %s\n", nodeName, familyKey)
		} else {
			index := indexOf(nodes, nodeName)
			productFamily[familyKey] = append(nodes[:index], nodes[index+1:]...)
			fmt.Printf("Node %s deleted from family %s\n", nodeName, familyKey)
		}
	} else {
		fmt.Printf("Family %s does not exist in the map\n", familyKey)
	}
}

func displayFamily(productFamily map[string][]string) {
	fmt.Println("Family -> Nodes")
	for key, values := range productFamily {
		fmt.Printf(" %s -> %v\n", key, values)
	}
}

func contains(slice []string, element string) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func indexOf(slice []string, element string) int {
	for i, v := range slice {
		if v == element {
			return i
		}
	}
	return -1
}

func main() {
	productFamily := map[string][]string{
		"7750": {"SR-1e", "SR-14S", "SR-12", "SR-7"},
		"7250": {"IXR-6", "IXR-R6", "IXR-R6dl"},
		"7705": {"SAR-8", "SAR-M", "SAR-A", "SAR-W"},
	}

	addNodeToFamily(productFamily)

	deleteNodeFromFamily(productFamily)

	addNewFamily(productFamily)

	displayFamily(productFamily)
}
