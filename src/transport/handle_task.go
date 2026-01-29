package transport

import "go-clean-app-example/src/domain/dto"

// HandleTask handles a Task request
// It takes a TransportRequest with dto.TaskRequest data and returns a TransportResponse with dto
func (h *BaseHandler) HandleTask(req TransportRequest[dto.TaskRequest]) TransportResponse[dto.TaskResponse] {
	// Log the request details
	// This is where you would typically log the request for debugging or monitoring purposes
	h.logger.Info("Handling Task request", map[string]interface{}{
		"source":      req.Source,
		"title":       req.Data.Title,
		"description": req.Data.Description,
	})

	// Call the use case with the request data
	result, err := h.useCases.CreateTask(req.Context, req.Data)

	// Handle errors and convert to TransportResponse
	if err != nil {
		return TransportResponse[dto.TaskResponse]{
			Success: false,
			Error:   err.Error(),
			Source:  req.Source,
		}
	}

	// Check if the result is successful and return the appropriate TransportResponse
	if result.Success {
		return TransportResponse[dto.TaskResponse]{
			Success: true,
			Data:    result.Data,
			Source:  req.Source,
		}
	}

	// If the result is not successful, return an error respons
	return TransportResponse[dto.TaskResponse]{
		Success: false,
		Error:   result.Error,
		Source:  req.Source,
	}
}
