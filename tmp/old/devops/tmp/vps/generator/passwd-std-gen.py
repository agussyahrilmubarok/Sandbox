import sys
import random
import string
import time
import argparse
from threading import Thread

# Default values
DEFAULT_LENGTH = 20
MIN_LENGTH = 8
SYMBOLS = "@#$%^&*()_+!~`|}{:;?><,./"
DEFAULT_OUTPUT_FILE = "password.txt"

def generate_password(length, symbols):
    """Generate a password with at least one symbol and one uppercase letter."""
    all_characters = string.ascii_letters + string.digits + symbols

    while True:
        password = ''.join(random.choice(all_characters) for _ in range(length))
        if any(char.isupper() for char in password) and any(char in symbols for char in password):
            return password

def show_loading(duration):
    """Show a loading animation for the given duration."""
    chars = '/-\\|'
    end_time = time.time() + duration
    while time.time() < end_time:
        for char in chars:
            sys.stdout.write(f'\r{char} Generating password...')
            sys.stdout.flush()
            time.sleep(0.1)

def main():
    """Main function to handle argument parsing and password generation."""
    parser = argparse.ArgumentParser(description='Generate a secure password with specified length.')
    parser.add_argument('--length', type=int, default=DEFAULT_LENGTH, help='Length of the password')
    parser.add_argument('--output_file', type=str, default=DEFAULT_OUTPUT_FILE, help='File to save the generated password')
    args = parser.parse_args()

    length = args.length
    output_file = args.output_file

    # Ensure the length is at least the minimum required
    if length < MIN_LENGTH:
        print(f"Length must be at least {MIN_LENGTH} characters.")
        sys.exit(1)

    # Record the start time
    start_time = time.time()

    # Start the loading animation in a separate thread
    loading_thread = Thread(target=show_loading, args=(5,))
    loading_thread.start()

    # Generate and save the password
    password = generate_password(length, SYMBOLS)
    with open(output_file, 'w') as f:
        f.write(f"Generated Password: {password}\n")

    # Wait for the loading animation to finish
    loading_thread.join()

    # Record the end time
    end_time = time.time()

    # Calculate elapsed time
    elapsed_time = end_time - start_time

    # Display the result
    print(f"\nPassword has been saved to {output_file}")
    with open(output_file, 'r') as f:
        print(f.read())

    # Display the time taken to generate the password
    print(f"Time taken to generate the password: {elapsed_time:.2f} seconds")

if __name__ == '__main__':
    main()
