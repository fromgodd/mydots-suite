import os
import subprocess

def get_mingw_version():
    mingw_path = "C:/MinGW"
    bin_path = os.path.join(mingw_path, "bin")
    executable = "gcc"  # You can choose any MinGW executable (e.g., gcc) for version check

    if os.path.exists(bin_path) and os.path.isdir(bin_path):
        try:
            result = subprocess.run(
                [os.path.join(bin_path, executable), "--version"],
                stdout=subprocess.PIPE,
                stderr=subprocess.PIPE,
                text=True,
                check=True
            )

            # Extract and print the version information
            version_lines = result.stdout.strip().splitlines()
            version = version_lines[0].split()[-1]
            print(f"{executable}, MinGW Version - {version} ✅")

        except subprocess.CalledProcessError as e:
            print(f"Error executing {executable} --version: {e.stderr}")

    else:
        print("MinGW is not installed in the default location.")

if __name__ == "__main__":
    get_mingw_version()
