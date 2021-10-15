package sec

type Query struct {
	DateRange  string   `json:"dateRange"`
	Category   string   `json:"category"`
	Ciks       []string `json:"ciks"`
	EntityName string   `json:"entityName"`
	Forms      []string `json:"forms"`
	Startdt    string   `json:"startdt"`
	Enddt      string   `json:"enddt"`
}

type Results struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore interface{}   `json:"max_score"`
		Hits     []interface{} `json:"hits"`
	} `json:"hits"`
	Aggregations struct {
		EntityFilter struct {
			DocCountErrorUpperBound int           `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int           `json:"sum_other_doc_count"`
			Buckets                 []interface{} `json:"buckets"`
		} `json:"entity_filter"`
		SicFilter struct {
			DocCountErrorUpperBound int           `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int           `json:"sum_other_doc_count"`
			Buckets                 []interface{} `json:"buckets"`
		} `json:"sic_filter"`
		BizStatesFilter struct {
			DocCountErrorUpperBound int           `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int           `json:"sum_other_doc_count"`
			Buckets                 []interface{} `json:"buckets"`
		} `json:"biz_states_filter"`
		FormFilter struct {
			DocCountErrorUpperBound int           `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int           `json:"sum_other_doc_count"`
			Buckets                 []interface{} `json:"buckets"`
		} `json:"form_filter"`
	} `json:"aggregations"`
	Query struct {
		Query struct {
			Bool struct {
				Must []struct {
					Match struct {
						DisplayNames string `json:"display_names"`
					} `json:"match"`
				} `json:"must"`
				MustNot []interface{} `json:"must_not"`
				Should  []interface{} `json:"should"`
				Filter  []struct {
					Terms struct {
						RootForm []string `json:"root_form"`
					} `json:"terms,omitempty"`
					Term struct {
						Sequence struct {
							Value int `json:"value"`
						} `json:"sequence"`
					} `json:"term,omitempty"`
					Range struct {
						FileDate struct {
							Gte string `json:"gte"`
							Lte string `json:"lte"`
						} `json:"file_date"`
					} `json:"range,omitempty"`
				} `json:"filter"`
			} `json:"bool"`
		} `json:"query"`
		From         int `json:"from"`
		Size         int `json:"size"`
		Aggregations struct {
			FormFilter struct {
				Terms struct {
					Field string `json:"field"`
					Size  int    `json:"size"`
				} `json:"terms"`
			} `json:"form_filter"`
			EntityFilter struct {
				Terms struct {
					Field string `json:"field"`
					Size  int    `json:"size"`
				} `json:"terms"`
			} `json:"entity_filter"`
			SicFilter struct {
				Terms struct {
					Field string `json:"field"`
					Size  int    `json:"size"`
				} `json:"terms"`
			} `json:"sic_filter"`
			BizStatesFilter struct {
				Terms struct {
					Field string `json:"field"`
					Size  int    `json:"size"`
				} `json:"terms"`
			} `json:"biz_states_filter"`
		} `json:"aggregations"`
		Sort []struct {
			FileDate struct {
				Order string `json:"order"`
			} `json:"file_date"`
		} `json:"sort"`
	} `json:"query"`
}


