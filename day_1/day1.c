#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdlib.h>

typedef struct Node {
    int data;
    struct Node* next;
}Node;

// Function to create a new node
Node* create_node(int data) {
    Node* new_node = (Node*)malloc(sizeof(Node)); // Allocate memory for a new node
    if (!new_node) {
        printf("Memory allocation failed.\n");
        return(NULL);
    }
    new_node->data = data;
    new_node->next = NULL;
    return new_node;
}

// Function to insert a new node into the ordered linked list
void ordered_insert(Node** head, int data) {
    Node* new_node = create_node(data);
    
    // If the list is empty or the new data should go before the first node
    if (*head == NULL || (*head)->data >= data) {
        new_node->next = *head;
        *head = new_node;
        return;
    }

    // Traverse the list to find the correct position for the new node
    Node* current = *head;
    while (current->next != NULL && current->next->data < data) {
        current = current->next;
    }

    // Insert the new node after the current node
    new_node->next = current->next;
    current->next = new_node;
}

// Function to insert a new node into the linked list
void insert(Node** head, int data) {
    Node* new_node = create_node(data);
    new_node->next = *head;
    *head = new_node;
}

// Function to print the linked list
void print_list(Node* head) {
    Node* current = head;
    while (current != NULL) {
        printf("%d -> ", current->data);
        current = current->next;
    }
    printf("NULL\n");
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

int count_first(Node* right, Node* left){
    int result = 0;

    while(right != NULL){
        result += abs(right->data - left->data);
        right = right->next;
        left = left->next;
    }
    return result;
}

int count_second(Node* left, Node* right){
    int result = 0;
    int quantity; // Amount of times the value in the left list apears in the right list
    Node* iteration;

    while(left != NULL){
        Node* iteration = right;
        quantity = 0;
        while(iteration != NULL){
            if(left->data == iteration->data){
                quantity ++;
            }
            iteration = iteration->next;
        }
        printf("%d\n", quantity);
        result += left->data * quantity;
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
    if (file == NULL) {
        printf("Error: Could not open file.\n");
        return 1;  // Exit if file could not be opened
    }

    Node* right = NULL;
    Node* left = NULL;
    
    char buffer[15];  // Buffer to hold the line, all lines have 15 characters
    int left_value = 0;
    int right_value = 0;
    int result = 0;

    switch (value){
        case 1:
            while (fgets(buffer, sizeof(buffer), file) != NULL) {
                if (sscanf(buffer, "%d   %d", &left_value, &right_value) != 2) {
                    // Error reading the input
                    printf("Error reading values.\n");
                    break;
                }
                ordered_insert(&left, left_value);
                ordered_insert(&right, right_value);
            }

            result = count_first(right, left);
            printf("Result is %d (clap)\n",result);

            free_list(&right);
            free_list(&left);
            break;
        case 2:
            while (fgets(buffer, sizeof(buffer), file) != NULL) {
                if (sscanf(buffer, "%d   %d", &left_value, &right_value) != 2) {
                    // Error reading the input
                    printf("Error reading values.\n");
                    break;
                }
                insert(&left, left_value);
                insert(&right, right_value);
            }

            result = count_second(left, right);
            printf("Result is %d (clap)\n",result);

            free_list(&right);
            free_list(&left);
            break;
        default:
            printf("Invalid number recieved\n");
            break;
    }

    fclose(file);  // Close the file

    printf("That's all folks!\n");
    return 0;
}