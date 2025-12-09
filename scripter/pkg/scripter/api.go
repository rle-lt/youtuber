package scripter

type OpenRouterRequest struct {
	Model    string           `json:"model"`
	Messages []RequestMessage `json:"messages"`
}

type OpenRouterResponse struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
	Model    string `json:"model"`
	Object   string `json:"object"`
	Created  int    `json:"created"`
	Choices  []struct {
		Logprobs           any    `json:"logprobs"`
		FinishReason       string `json:"finish_reason"`
		NativeFinishReason string `json:"native_finish_reason"`
		Index              int    `json:"index"`
		Message            struct {
			Role             string `json:"role"`
			Content          string `json:"content"`
			Refusal          any    `json:"refusal"`
			Reasoning        string `json:"reasoning"`
			ReasoningDetails []struct {
				Format string `json:"format"`
				Index  int    `json:"index"`
				Type   string `json:"type"`
				Text   string `json:"text"`
			} `json:"reasoning_details"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens        int  `json:"prompt_tokens"`
		CompletionTokens    int  `json:"completion_tokens"`
		TotalTokens         int  `json:"total_tokens"`
		Cost                int  `json:"cost"`
		IsByok              bool `json:"is_byok"`
		PromptTokensDetails struct {
			CachedTokens int `json:"cached_tokens"`
			AudioTokens  int `json:"audio_tokens"`
			VideoTokens  int `json:"video_tokens"`
		} `json:"prompt_tokens_details"`
		CostDetails struct {
			UpstreamInferenceCost            any `json:"upstream_inference_cost"`
			UpstreamInferencePromptCost      int `json:"upstream_inference_prompt_cost"`
			UpstreamInferenceCompletionsCost int `json:"upstream_inference_completions_cost"`
		} `json:"cost_details"`
		CompletionTokensDetails struct {
			ReasoningTokens int `json:"reasoning_tokens"`
			ImageTokens     int `json:"image_tokens"`
		} `json:"completion_tokens_details"`
	} `json:"usage"`
}
