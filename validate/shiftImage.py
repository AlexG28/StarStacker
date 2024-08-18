from PIL import Image


def shift_image(image_path, output_path, x_shift, y_shift): 
    with Image.open(image_path) as img: 
        width, height = img.size

        shifted = Image.new(img.mode, (width, height))

        shifted.paste(img, (x_shift, y_shift))

        left = 0
        right = left + width

        upper = 0
        lower = upper + height


        if x_shift >= 0: 
            left += x_shift
        else: 
            right += x_shift


        if y_shift >= 0: 
            upper += y_shift
        else: 
            lower += y_shift

        shifted = shifted.crop((left, upper, right, lower))

        shifted.save(output_path)
        print(shifted.size)
        print(f"image saved to: {output_path}")



if __name__ == "__main__": 
    x_shift = 1
    y_shift = -9
    image_path = "/home/alexlinux/projects/StarCounter/testfiles/example.png"
    save_path = f"/home/alexlinux/projects/StarCounter/testfiles/exampleshifted_{x_shift}_{y_shift}.png"

    shift_image(image_path, save_path, x_shift, y_shift)
