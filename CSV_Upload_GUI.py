import tkinter as tk
from tkinter import filedialog, messagebox

class FileUploaderApp:
    def __init__(self, root):
        self.root = root
        self.root.title("File Uploader")
        self.files = []

        # Create a frame for buttons
        self.frame = tk.Frame(root)
        self.frame.pack(pady=20)

        # Button to add files
        self.add_file_button = tk.Button(self.frame, text="Add File", command=self.add_file)
        self.add_file_button.pack(side=tk.LEFT, padx=10)

        # Submit button
        self.submit_button = tk.Button(self.frame, text="Submit", command=self.submit_files)
        self.submit_button.pack(side=tk.LEFT, padx=10)

        # Label to display the uploaded files
        self.files_label = tk.Label(root, text="No files uploaded", justify=tk.LEFT)
        self.files_label.pack(pady=20)

    def add_file(self):
        """Opens file dialog to select files and adds them to the list."""
        file_path = filedialog.askopenfilename()
        if file_path:
            self.files.append(file_path)
            self.update_file_list()

    def update_file_list(self):
        """Updates the label to display the list of selected files."""
        if self.files:
            file_list = "\n".join(self.files)
            self.files_label.config(text=f"Uploaded Files:\n{file_list}")
        else:
            self.files_label.config(text="No files uploaded")

    def submit_files(self):
        """Reads the content of the files and stores them in a dictionary."""
        if not self.files:
            messagebox.showwarning("Warning", "No files uploaded")
            return

        file_contents = {}
        try:
            for file_path in self.files:
                with open(file_path, 'r') as file:
                    file_contents[file_path] = file.read()
            
            # Display success message
            messagebox.showinfo("Success", f"Successfully read {len(self.files)} files.")
            
            # Do something with the file contents (store in variables, etc.)
            # For demonstration, we print the contents
            for file_name, content in file_contents.items():
                print(f"File: {file_name}")
                print(content)
                print("-" * 50)
                
        except Exception as e:
            messagebox.showerror("Error", f"Error reading files: {str(e)}")

if __name__ == "__main__":
    root = tk.Tk()
    app = FileUploaderApp(root)
    root.mainloop()
