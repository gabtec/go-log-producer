package utils

func GetHttpCode(idx int) int {
	possibleResults := []int{200, 400, 401, 403, 404, 500}
	return possibleResults[idx]
}

func GetHttpErrorCode(idx int) int {
	possibleResults := []int{404, 400, 401, 403, 404, 500}
	return possibleResults[idx]
}

func GetErrorMessage(code int) string {
	messages := map[int]string{
		400: "Account not found",
		401: "You are not authorized to create users",
		403: "You have no permissions to delete accounts",
		404: "Missing required field: 'X'",
		500: "Unexpected server behavior",
	}

	return messages[code]
}

func GetSuccessMessage(idx int) string {
	messages := []string{
		"Object created successfully",
		"Object updated successfully",
		"Object deleted successfully",
		"User X loggedIn successfully",
		"Cronjob executed successfully",
	}

	return messages[idx]
}
