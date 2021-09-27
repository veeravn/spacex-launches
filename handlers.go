package main

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/machinebox/graphql"
)

func getLaunches(c echo.Context) error {

	// In case of any error, we respond with an error to the user
	limit := "10"
	offset := "0"

	if c.QueryParam("limit") != "" {
		limit = c.QueryParam("limit")
	}
	if c.QueryParam("offset") != "" {
		offset = c.QueryParam("offset")
	}
	query := `
    {
  launchesPast(limit: ` + limit + `, offset: ` + offset + `) {
  mission_name
  launch_date_local
  launch_site {
  site_name_long
  }
  links {
  article_link
  video_link
  }
  rocket {
  rocket_name
  first_stage {
    cores {
      flight
      core {
        reuse_count
        status
      }
    }
  }
  second_stage {
    payloads {
      payload_type
      payload_mass_kg
      payload_mass_lbs
    }
  }
  }
  ships {
  name
  home_port
  image
  }
  }
  }
    `

	//create graphql client
	graphqlClient := graphql.NewClient("https://api.spacex.land/graphql/")
	graphqlRequest := graphql.NewRequest(query)
	ctx := context.Background()
	var resp = Data{}

	if err := graphqlClient.Run(ctx, graphqlRequest, &resp); err != nil {
		panic(err)
	}
	var launch LaunchesPastTable
	var launches []LaunchesPastTable
	for _, key := range resp.LaunchesPast {
		launch = serializeJSON(key)
		launches = append(launches, launch)
		// if i == 0 {
		// 	c.Request.Write([]byte("["))
		// }
		// bytes := serializeJSON(key)
		// c.Write(bytes)
		// if i == len(resp.LaunchesPast)-1 {
		// 	c.Write([]byte("]"))
		// } else {
		// 	c.Write([]byte(","))
		// }
	}
	return c.JSON(http.StatusOK, launches)
	// http.Redirect(w, r, "/assets/", http.StatusOK)
}

func serializeJSON(key LaunchesPast) LaunchesPastTable {
	launch := LaunchesPastTable{
		SiteNameLong: key.LaunchSite.SiteNameLong,
		MissionName:  key.MissionName,
		RocketName:   key.Rocket.RocketName,
	}
	return launch
}
