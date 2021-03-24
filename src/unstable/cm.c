#include <iostream>
#include <Windows.h>

int main(int argc,char *argv[]){
	::ShowWindow(::GetConsoleWindow(), SW_SHOW);
	std::system((std::string("cd picocrypt && start.vbs \"")+argv[1]+"\"").c_str());
}