#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdlib.h>

typedef struct Node {
    struct Node* next;
    int size;
    int data[]; // Flexible array, never used it before. 
    // In order for it to work it needs to be declared as the last element of the struct. 
    // Also, when alocating memory for the node, i need to add the memory for the array as well.
}Node;

// Function to create a new node
Node* create_node(int data[], int size) {
    Node* new_node = (Node*)malloc(sizeof(Node) + size * sizeof(int)); // Allocate memory for a new node
    if (!new_node) {
        printf("Memory allocation failed.\n");
        return(NULL);
    }
    new_node->size = size;
    new_node->next = NULL;

    // Copy the data into the flexible array
    for (int i = 0; i < size; i++) {
        new_node->data[i] = data[i];
    }

    return new_node;
}

// Function to insert a new node into the linked list
void insert(Node** head, int data[], int size) {
    Node* new_node = create_node(data, size);
    new_node->next = *head;
    *head = new_node;
}

// Function to print a node
void print_node(Node* node) {
    if (node == NULL) {
        printf("NULL");
    }else{
        printf("Data: [");
        for (int i = 0; i < node->size; i++) {
            printf("%d-", node->data[i]);
        }
    
    }
    printf("], ");
}

// Function to print the linked list
void print_list(Node* head) {
    Node* current = head;
    while (current != NULL) {
        print_node(current);
        current = current->next;
    }
}

// Function to free all nodes in the linked list
void free_list(Node** head) {
    Node* current = *head;
    Node* next_node = NULL;
    
    while (current != NULL) {
        next_node = current->next;   // Save the next node
        free(current);               // Free the current node
        current = next_node;         // Move to the next node
    }
    
    *head = NULL;  // Set head to NULL to indicate the list is empty
}

size_t get_line_size(FILE* file) {
    size_t size = 0;
    int ch;
    // Move the file pointer to the beginning of the line
    while ((ch = fgetc(file)) != EOF && ch != '\n') {
        size++;
    }
    // Rewind the file pointer to the beginning of the line
    fseek(file, -size - 1, SEEK_CUR);
    return size;
}

char* read_line(FILE* file) {
    // First, determine the size of the line
    size_t line_size = get_line_size(file);
    if (line_size == 0) return NULL;

    // Allocate memory for the line plus the null terminator
    char* line = malloc(line_size + 1); // +1 for null terminator
    if (!line) {
        perror("Memory allocation failed");
        return NULL;
    }

    // Now, read the line into the allocated memory
    size_t i = 0;
    int ch;
    while (i < line_size && (ch = fgetc(file)) != EOF && ch != '\n') {
        line[i++] = (char)ch;
    }

    line[i] = '\0'; // Null-terminate the string
    return line;
}

int count_first(Node* right, Node* left){
    int result = 0;

    while(right != NULL){
        result += abs(right->data - left->data);
        right = right->next;
        left = left->next;
    }
    return result;
}

int main(int argc, char *argv[]) {
    printf("Starting Program!\n");

    // Check if the correct number of arguments is provided
    if (argc < 2) {
        printf("Usage: %s <input_file> <query_number>\n", argv[0]);
        printf("Valid numbers for querys are 1 and 2\n");
        return 1;
    }

    // Convert command-line arguments to integers
    int value = atoi(argv[2]);

    // Access input file
    FILE *file = fopen(argv[1], "r");  // Open file in read mode
    if (!file) {
        printf("Error: Could not open file.\n");
        return 1;  // Exit if file could not be opened
    }

    Node* data = NULL;
    
    int size = 0;
    int result = 0;
    char* line = NULL;
    size_t count = 0;

    switch (value){
        case 1:

            while ((line = read_line(file)) != NULL) {
                // Count how many numbers there are (split by spaces)
                count = 0;
                char* token = strtok(line, " ");
                while (token != NULL) {
                    count++;
                    token = strtok(NULL, " ");
                }

                // Allocate memory for the integer array
                int numbers[count];
                if (numbers == NULL) {
                    perror("Memory allocation failed");
                    return 0;
                }

                // Parse the numbers into the array
                token = strtok(line, " "); // Reset strtok to the original string
                size_t i = 0;
                while (token != NULL) {
                    numbers[i++] = atoi(token); // Convert token to an integer and store it
                    token = strtok(NULL, " ");
                }

                char* token = strtok(str, " ");
                int* numbers = NULL;

                while (token != NULL) {
                    // Allocate memory for the integer array only when needed
                    numbers = (int*)realloc(numbers, (count + 1) * sizeof(int));
                    if (numbers == NULL) {
                        perror("Memory allocation failed");
                        return NULL;
                    }

                    // Convert the token to an integer and store it
                    numbers[count] = atoi(token);
                    count++;

                    // Get the next token
                    token = strtok(NULL, " ");
                }

                *num_elements = count; // Set the number of elements

                insert(&data,numbers,count);
                free(line); // Free the memory after processing the line
            }
            
            print_list(data);

            result = 0; // NEED TO CALCULATE THE RESULT
            printf("There are %d safe registers (clap)\n",result);

            free_list(&data);
            break;
        case 2:
            break;
        default:
            printf("Invalid number recieved\n");
            break;
    }

    fclose(file);  // Close the file

    printf("That's all folks!\n");
    return 0;
}