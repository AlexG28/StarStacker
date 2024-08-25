from PIL import Image
import numpy as np

def add_noise(image) -> Image: 
    img_array = np.array(image)

    mean = 20
    var = 10
    sigma = var ** 0.5 

    noise = np.random.normal(mean, sigma, img_array.shape)

    noisy_img_arr = img_array + noise 

    noisy_img_arr = np.clip(noisy_img_arr, 0, 255)
    noisy_img_arr = noisy_img_arr.astype(np.uint8)

    return Image.fromarray(noisy_img_arr)


def shift_image(image_path, x_shift, y_shift): 
    with Image.open(image_path) as img: 
        width, height = img.size
        corrected = Image.new("RGB", img.size, (0, 0, 0))
        shifted = Image.new(img.mode, (width, height))

        shifted.paste(img, (x_shift, y_shift))
        
        corrected.paste(shifted, mask=shifted.split()[3])

        corrected = add_noise(corrected)
        return corrected
        

def main(): 
    x_shift = 5
    y_shift = 5
    number_of_images = 1
    image_path = "/home/alexlinux/projects/StarCounter/testfiles/6stars.png"
    image_output_dir = "/home/alexlinux/projects/StarCounter/testfiles"

    for i in range(number_of_images): 
        image = shift_image(image_path, x_shift, y_shift)

        image_name = f"{image_output_dir}/image{i}_{x_shift}_{y_shift}.png" 
        image.save(image_name)


if __name__ == "__main__": 
    main()