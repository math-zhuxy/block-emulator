go build -o blockEmulator_Windows_Precompile.exe main.go

start cmd /k blockEmulator_Windows_Precompile.exe -n 0 -N 4 -s 0 -S 1 & 

start cmd /k blockEmulator_Windows_Precompile.exe -n 1 -N 4 -s 0 -S 1 & 

start cmd /k blockEmulator_Windows_Precompile.exe -n 2 -N 4 -s 0 -S 1 & 

start cmd /k blockEmulator_Windows_Precompile.exe -n 3 -N 4 -s 0 -S 1 & 

start cmd /k blockEmulator_Windows_Precompile.exe -c -N 4 -S 1 & 

