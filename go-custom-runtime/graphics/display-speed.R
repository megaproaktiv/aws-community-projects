speedIn <- read.csv("./speed.csv", sep = ";", header = TRUE)
speed <- transform(speedIn, 
                  runtime = as.character(Runtime), 
                  init=(as.numeric(Init)),
                  cold=(as.numeric(Init)+as.numeric(Cold))
                  )

library(ggplot2)
ggplot(speed, aes(y= init, x=runtime)) + geom_boxplot()


ggplot(speed, aes(y= cold, x=runtime)) + geom_boxplot()
