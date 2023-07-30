speedIn <- read.csv("./speed-x86-128.csv", sep = ";", header = TRUE)
speed <- transform(speedIn, 
                  Name = as.character(Name), 
                  init=(as.numeric(Init)),
                  cold=(as.numeric(Cold)),
                  sum=(as.numeric(Init)+as.numeric(Cold))
                  )

library(ggplot2)
ggplot(speed, aes(y= init, x=Name)) + geom_boxplot() +
  ylim(0, NA) 


ggplot(speed, aes(y= cold, x=Name)) + geom_boxplot()


ggplot(speed, aes(y= sum, x=Name)) + geom_boxplot()

