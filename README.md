1. Get the exchange rate API key link: https://www.exchangerate-api.com/ just click on get free key
2. setup a .env file with `EXCHANGE_RATE_API_KEY` env variable in the root directory
3. run go api using `go run .` make sure you are in the root directory
4. in another terminal, cd into `react-app` directory, install packages using `npm i` and start app using `npm run start`
5. Select `from`, `to` and enter the `amount`. Afterwards, click on `convert` and volla, you should have the exchange rate. 


Important notes: 
1. Make sure golang API is working on port 8080. If something is running on 8080 then first kill it before running the go API. React app listens to 8080 for the golang API.
2. The app always have significant room of improvement for being production ready. This is merely an exercise completed for this interview
3. This was fun to build