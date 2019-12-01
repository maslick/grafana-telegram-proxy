package main

type Request struct {
	Metrics  []EvalMatch `json:"evalMatches"`
	Mess     string      `json:"message"`
	RuleName string      `json:"ruleName"`
	RuleUrl  string      `json:"ruleUrl"`
	State    string      `json:"state"`
	Title    string      `json:"title"`
}

type EvalMatch struct {
	Value  float32           `json:"value"`
	Metric string            `json:"metric"`
	Tags   map[string]string `json:"tags"`
}

type Message struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}
