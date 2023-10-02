#include <iostream>
#include <string>

int main() {
    std::string user_input;
    std::cout << "Entrez le flag : ";
    std::cin >> user_input;

    if (user_input == "Flag123") {
        std::cout << "Bravo, vous avez trouvé le drapeau !" << std::endl;
    } else {
        std::cout << "Désolé, essayez à nouveau." << std::endl;
    }

    return 0;
}
