package sec

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	searchUrl = "https://efts.sec.gov/LATEST/search-index"
)

func dataSourceSecS1Filing() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSecS1FilingRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceSecS1FilingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var results Results
	currentTime := time.Now().Format("2006-01-02")

	name := d.Get("name").(string)

	query := Query{
		EntityName: name,
		DateRange:  "all",
		Category:   "custom",
		Forms:      []string{"S-1"},
		Startdt:    "2001-01-01",
		Enddt:      currentTime,
	}

	postBody, _ := json.Marshal(query)

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(searchUrl, "application/json", responseBody)
	if err != nil {
		diag.FromErr(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		diag.FromErr(err)
	}

	json.Unmarshal([]byte(body), &results)

	d.SetId(name)

	if results.Hits.Total.Value > 0 {
		d.Set("status", true)
	} else {
		d.Set("status", false)
	}
	return nil
}
