from PIL import Image 


def shift_image(image_path, output_path, x, y): 
    with Image.open(image_path) as img: 
        width, height = img.size

        shifted = Image.new(img.mode, (width, height))

        shifted.paste(img, (x, y))

        shifted.save(output_path)
        print(f"image saved to: {output_path}")



if __name__ == "__main__": 
    image_path = "/home/alexlinux/projects/StarCounter/testfiles/example2.png"
    save_path = "/home/alexlinux/projects/StarCounter/testfiles/example2shifted.png"
    x_shift = 10
    y_shift = 20

    shift_image(image_path, save_path, x_shift, y_shift)
