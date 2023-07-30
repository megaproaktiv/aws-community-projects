library(readr)
library(ggplot2)

df <- read.csv2("benchmarks.csv")
bp <- barplot(t(df[ , -1]), col = c("blue", "red", "green", "orange", "gold"))
axis(side = 1, at = bp, labels = df$type)

ggplot(data=df, aes(x=type, y=compressing+decompressing, fill=type)) +
  geom_bar(stat="identity", position=position_dodge())+
  geom_text(aes(label=type), vjust=1.6, color="white",
            position = position_dodge(0.9), size=3.5)+
  scale_fill_brewer(palette="Paired")+
  theme_classic()
