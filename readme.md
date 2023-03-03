eu to voando alto

sigo na minha luta e trabalhando dobrado

eu to voando alto

sigo na minha luta e trabalhando dobrado



```go
    for day := range weekdays {
        _ := os.Open("window")
        
        routes.TurnOnPC("localhost")
        routes.HaveBreakfast()
        
        state := models.NewDay {
            day: day, 
            activeMood: 6, // neutral
            // ...
        }

        // 8 to 12
        for () {
            load, stress := models.Work.Start()
            if load != nil {
                if load.size == "low" { 
                    load.ToneDownWork()
                    go func() async {
                        await models.Spotify("Yung Buda", "Hacker Dresscode").play()
                        await models.Spotify("Luv Resval", "Tout s'en va").play()
                        await models.Spotify("Nelson Gonçalves", "Naquela Mesa").play()
                        await models.Spotify("The Killers", "Somebody Told Me").play()
                        await models.Spotify("Mariana Froes", "Mais Uma Canção Para Você").play()
                    }
                }
                load.ToneUpWork()
            }

            if load == nil {
                load.Pause()
                study, err := models.Study("Flutter");
                if err == nil {
                    state.active = state.active + 2
                } else {
                    isLoadAvailable := load.Check() // implement
                    if !isLoadAvailable {
                        research, errResearch := study.research()
                        if errResearch != nil {
                            state.active = state.active - 1
                            load.Resume("idle")
                        }
                    }
                }
                for () {

                }
                
            }



        }

        

        studyGolang();
        studyFlutter();

        switch day {
            case "Monday", "Wednesday":
                GoOut("Volley", crush)
            case "Friday":
                models.Spotify(
                    "Grupo Menos É Mais",
                    "Pout-Pourri: Melhor Eu Ir
                    /Ligando os Fatos
                    /Sonho de Amor
                    /Deixa Eu Te Querer
                    - Ao Vivo"
                ).play
        }

    }
```
