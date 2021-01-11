#!/usr/bin/env bash

  time_val=$(date)                                                               
  echo -ne "\\r$time_val"                                                        
                                                                                  
  time_min=$(echo $time_val | cut -d' ' -f 5)                                    
  time_min=${time_min:3:5}                                                       
echo $time_min
grep '10' <(echo $time_min)                                      
echo end 
