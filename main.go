package main

import (
	"fmt"

	"gitlab.com/snowmerak/simple-html-generator/e"
)

func main() {
	titleString := "Serve it Yourself!"

	metas := e.Elements{
		e.Text{
			Value: "<meta charset=\"UTF-8\">",
		},
		e.Text{
			Value: "<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">",
		},
		e.Text{
			Value: "<meta http-equiv=\"X-UA-Compatible\" content=\"ie=edge\">",
		},
		e.Text{
			Value: "<title>" + titleString + "</title>",
		},
	}

	title := e.Elements{
		e.Header{
			Level: 1,
			Value: e.Text{
				Value: titleString,
			},
		},
		e.Bold{
			Value: e.Italic{
				Value: e.Text{
					Value: "\"Serve it Yourself!\"는 백엔드 인프라를 직접 구현해보는 프로젝트입니다.",
				},
			},
		},
	}

	summery := e.Elements{
		e.Header{
			Level: 2,
			Value: e.Text{
				Value: "개요",
			},
		},
		e.Paragraph{
			Value: e.Text{
				Value: "",
			},
		},
	}

	roadmap := e.Elements{
		e.Header{
			Level: 2,
			Value: e.Text{
				Value: "구상도",
			},
		},
		e.Text{
			Value: "<figure><img src=\"out/uml/roadmap/roadmap.svg\" alt=\"roadmap\"></figure>",
		},
	}

	contributer := e.Elements{
		e.Header{
			Level: 2,
			Value: e.Text{
				Value: "기여자",
			},
		},
		e.Paragraph{
			Value: e.Text{
				Value: "snowmerak: <a href=\"https://gitlab.com/snowmerak\">Gitlab</a>",
			},
		},
	}

	communityDatas := [][]string{
		{"코딩맛집", "<a href=\"https://discord.gg/U7HhCCzV\">discord</a>"},
		{"위클리 아카데미", "<a href=\"https://weekly.ac\">homepage</a>"},
	}
	communityRows := []e.TableRow{}
	for _, communityData := range communityDatas {
		row := e.TableRow{}
		for _, value := range communityData {
			row.Values = append(row.Values, e.TableData{
				Values: []e.Element{
					e.Text{
						Value: value,
					},
				},
			})
		}
		communityRows = append(communityRows, row)
	}
	community := e.Elements{
		e.Header{
			Level: 2,
			Value: e.Text{
				Value: "커뮤니티",
			},
		},
		e.Table{
			Headers: e.TableHeader{
				Values: []e.Element{
					e.Text{
						Value: "이름",
					},
					e.Text{
						Value: "링크",
					},
				},
			},
			Rows: communityRows,
		},
	}

	projectDatas := [][]string{
		{"Lux", "snowmerak", "simple webframework for golang", "<a href=\"https://gitlab.com/serve-it-yourself/lux\">Gitlab</a>"},
		{"logstream", "snowmerak", "log library for golang", "<a href=\"https://gitlab.com/serve-it-yourself/logstream\">Gitlab</a>"},
		{"log-silo", "snowmerak", "simple log store based parquet", "<a href=\"https://gitlab.com/serve-it-yourself/log-silo\">Gitlab</a>"},
		{"compositor", "snowmerak", "reverse proxy server based docker", "<a href=\"https://gitlab.com/serve-it-yourself/compositor\">Gitlab</a>"},
		{"virtual-gate", "snowmerak", "a load balancer and rate limitter", "<a href=\"https://gitlab.com/serve-it-yourself/virtual-gate\">Gitlab</a>"},
	}
	projectsRows := []e.TableRow{}
	for _, projectData := range projectDatas {
		row := e.TableRow{}
		for _, value := range projectData {
			row.Values = append(row.Values, e.TableData{
				Values: []e.Element{
					e.Text{
						Value: value,
					},
				},
			})
		}
		projectsRows = append(projectsRows, row)
	}
	projects := e.Elements{
		e.Header{
			Level: 2,
			Value: e.Text{
				Value: "프로젝트",
			},
		},
		e.Table{
			Headers: e.TableHeader{
				Values: e.Elements{
					e.Text{
						Value: "프로젝트명",
					},
					e.Text{
						Value: "메인테이너",
					},
					e.Text{
						Value: "개요",
					},
					e.Text{
						Value: "링크",
					},
				},
			},
			Rows: projectsRows,
		},
	}

	doc := e.Document{
		Head: e.Head{
			Values: []e.Element{
				metas,
				e.Link{
					Href: "./mvp.css",
				},
			},
		},
		Body: e.Body{
			Values: []e.Element{
				title,
				e.Hr{},
				summery,
				e.Hr{},
				roadmap,
				e.Hr{},
				contributer,
				e.Hr{},
				community,
				e.Hr{},
				projects,
				e.Hr{},
			},
		},
	}
	fmt.Println(doc.Parse())
}
