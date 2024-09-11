package api

import (
	"mawa3id/DB"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const apiKey = "a5fdf71cd14000eadb6ba3f508ed46a7"

type LocationResponse struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func SearchServicesBySubCategory(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		subcategory := ctx.Query("q")
		byRatingStr := ctx.Query("r")
		byRating := false
		if byRatingStr == "true" {
			byRating = true
		}
		// myPlaceStr := ctx.Query("place")

		// url := fmt.Sprintf("http://api.positionstack.com/v1/forward?access_key=%s&query=%s", apiKey, myPlaceStr)

		// // Make the GET request to the LocationIQ API
		// resp, err := http.Get(url)
		// if err != nil {
		// 	log.Fatalf("Error making the API request: %v", err)
		// }
		// defer resp.Body.Close()

		// // Read the response body
		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	log.Fatalf("Error reading the response body: %v", err)
		// }
		// return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		// 	"ok":  true,
		// 	"msg": string(body),
		// })
		// // Parse the JSON response
		// var locationData []LocationResponse
		// err = json.Unmarshal(body, &locationData)
		// if err != nil {
		// 	log.Fatalf("Error parsing the API response: %v", err)
		// }

		// // Check if any location was found
		// longitude := ""
		// latitude := ""
		// if len(locationData) > 0 {
		// 	// Print the first location's details (latitude, longitude, and display name)
		// 	fmt.Printf("Latitude: %s\n", locationData[0].Latitude)
		// 	fmt.Printf("Longitude: %s\n", locationData[0].Longitude)
		// 	longitude = locationData[0].Longitude
		// 	latitude = locationData[0].Latitude
		// } else {
		// 	fmt.Println("No location found")
		// }
		// // myPlace := strings.ReplaceAll(myPlaceStr, " ", "")
		// // longitude := strings.Split(myPlace, ",")[1]
		// // // remove space from the latitude
		// // latitude := strings.Split(myPlace, ",")[0]
		// // //remove all the spaces from the latitude if it has any
		// // fmt.Println(longitude, latitude)
		// return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		// 	"ok":   true,
		// 	"msg":  "not implemented yet",
		// 	"long": longitude,
		// 	"lat":  latitude,
		// 	"q":    subcategory,
		// 	"r":    byRating,
		// })
		if subcategory != "" {
			subcategoryID, err := strconv.ParseInt(subcategory, 10, 32)
			if err != nil || subcategoryID == 0 {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": err,
					"msg":   "subcategory not found",
				})
			}
			services := []DB.Service{}
			switch byRating {
			case true:
				services, err = db.OrderServicesByRating(ctx.Context(), DB.OrderServicesByRatingParams{
					SubcategoryID: int32(subcategoryID),
					Limit:         25,
					Offset:        0})
				if err != nil {
					return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"ok":    false,
						"error": err,
					})
				}
			case false:
				services, err = db.SearchServicesBySubCategory(ctx.Context(), DB.SearchServicesBySubCategoryParams{
					SubcategoryID: int32(subcategoryID),
					Limit:         25,
					Offset:        0})
				if err != nil {
					return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"ok":    false,
						"error": err,
					})
				}
			}
			if len(services) == 0 {
				services = []DB.Service{}
			}
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"ok":       true,
				"services": services,
				"filter":   byRating,
			})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":    false,
			"error": "category or subcategory not found",
		})
	}
}
