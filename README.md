# Historic cars command line client 
---------------------------------------

  * [Intro](#intro)
  * [Requirements](#requirements)
  * [Usage](#usage)
    * [Option '1' : display full list of cars](display-full-list-of-cars) 
    * [Option '2' : search cars](search-cars) 
    * [Option '3' : display car statistics](display-car-statistics) 
* [License](#license)

---------------------------------------

## Intro

It is a command line client, implemented in Golang, enabling to process and extract statistics about car models produced in USA, Japan and Europe between 1978 and 1982, visualising some interesting trends. The origin of the data is [data.world](https://data.world/dataman-udit/cars-data) website.

## Requirements

Go 1.18 or higher

## Usage

To run the command line client

```
go run cars.go  
```

It displays a menu and waits for your choice of action 

```
= = = = = = = = = MAIN MENU = = = = = = = = = = = =

 0 => quit!

 1 => display full list of cars
 2 => search cars
 3 => display car statistics for US, European and Japanese cars (1970 - 1982)

= = = = = = = = = = = = = = = = = = = = = = = = = =

Your choice => 

```
### Option '1' : display full list of cars

Prints out the entire database of car models. At the time of writing, ca 400 items.

### Option '2' : search cars

You can search specific cars by name (e.g. Ford), model (e.g. Ranger) or country of origin (USA, Japan, Europe).

### Option '3' : display car statistics

The option generates 7 different bar charts and opens them on your default browser. 

p.s. if, for some reason, it does not open the browser, the generated file is located on "./temp/bar.html"

A sample bar chart...

![Sample bar chart](/data/avg-cylinders.png "Sample bar chart")

## License

[MIT](https://opensource.org/license/mit/)
