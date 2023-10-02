#include <iostream>
#include <string>
#include <cstdlib> // Pour accéder aux variables d'environnement
#include <fstream>
#include <sstream>
#include <iomanip>
#include <string>
// Fonction pour encoder en base64
std::string base64Encode(const std::string &input) {
    std::stringstream ss;
    for (size_t i = 0; i < input.size(); i += 3) {
        uint32_t val = (static_cast<uint32_t>(input[i]) << 16) |
                       (i + 1 < input.size() ? (static_cast<uint32_t>(input[i + 1]) << 8) : 0) |
                       (i + 2 < input.size() ? static_cast<uint32_t>(input[i + 2]) : 0);

        ss << "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"[(val >> 18) & 0x3F];
        ss << "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"[(val >> 12) & 0x3F];
        ss << "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"[(val >> 6) & 0x3F];
        ss << "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"[val & 0x3F];
    }
    return ss.str();
}

int main() {
    std::string secret_password = std::getenv("SECRETPASS"); // Obtenir le mot de passe depuis la variable d'environnement
    std::string  vi= "ZTFjNDViM2U3NjI3OGU2YjhlMzcwODM4NzhiOTExODc=";
    if (secret_password.empty()) {
        std::cout << "La variable d'environnement SECRETPASS n'est pas définie." << std::endl;
        return 1;
    }
    
    std::string user_input_base64;

    std::cout << "Entrez le drapeau (encodé en base64) : ";
    std::cin >> user_input_base64;

    // Décodez l'entrée utilisateur depuis base64
    std::string user_input;
    // Décodez l'entrée utilisateur depuis base64
    std::string user_input;
    for (size_t i = 0; i < user_input_base64.size(); i++) {
        if (user_input_base64[i] == '=') {
            break; // Ignore les caractères de padding
        }
        uint8_t val = 0;
        char c = user_input_base64[i];
        if (c >= 'A' && c <= 'Z') {
            val = c - 'A';
        } else if (c >= 'a' && c <= 'z') {
            val = c - 'a' + 26;
        } else if (c >= '0' && c <= '9') {
            val = c - '0' + 52;
        } else if (c == '+') {
            val = 62;
        } else if (c == '/') {
            val = 63;
        }
        user_input.push_back(static_cast<char>(val));
    }

    // Déchiffrement simple : décaler chaque caractère de 1 vers l'arrière dans l'alphabet.
    for (char &c : user_input) {
        if (std::isalpha(c)) {
            c = static_cast<char>(c - 1);
        }
    }

    if (user_input == secret_password) {
        std::cout << "Bravo, vous avez trouvé le drapeau !" << std::endl;
    } else {
        std::cout << "Désolé, essayez à nouveau." << std::endl;
    }

    return 0;
}
