# What does LazyParam do?

It is an automation tool for fuzzing on url(s) which is a combination of open source tools and a few small innovations.
The lazyparam script takes one or more urls from you, extracts the parameters and uses a custom wordlist of parameters too, then use x8 tool to param fuzz them as GET and POST.

## Features
- Work on one or more urls via input or file
- extract url parameters using open source tools
- Combine extracted params with a custom wordlist of params
- Configurable structure

## How to install and run
### Direct Download
 - You can download the executable file of the lazyparam from [**this link**](https://github.com/mahyarkermani1/LazyParam/releases/download/v0.1.0/lazyparam)

### Build from source
  - **Download Go**
    ```
    apt install golang-go
    ```
  
  - **Download repository**
    ```
    git clone https://github.com/mahyarkermani1/LazyParam.git
    ```
  
  - **cd to the project directory**
    ```
    cd LazyParam
    ```
  
  - **Install prerequisites**
    ```
    go mod tidy
    ```
  
  - Build and run
    ```
    go build
    ./lazyparam -u https://example.com/home
    ```

## Roadmap

- [x] Version 0.0.1 <details>
  <summary></summary>

  - [x] Designing the readme file
  - [x] Adding the project to GitHub
  - [x] Designing the feature list

</details>


- [x] Version 0.0.2 <details>
  <summary></summary>

  - [x] Designing the project tree
  - [x] Adding project files

</details>

- [x] Version 0.1.0 <details>
  <summary></summary>

   - [x] Parameter Fuzz on a single url
   - [x] Run the [fallparams ](https://github.com/ImAyrix/fallparams/)tool on the url to collect page parameters, combine them with a [custom wordlist](https://github.com/mahyarkermani1/LazyParam/blob/main/wordlists/custom_params.txt) of parameters, and perform parameter fuzz with [x8](https://github.com/Sh1Yo/x8)

</details>


## Open source tools used
- To extract page parameters
   - [**fallparams**](https://github.com/ImAyrix/fallparams)

- For parameter fuzzing
   - [**x8**](https://github.com/Sh1Yo/x8)

## Image gallery
Coming soon
