# coding: utf-8

import numpy as np
import PIL
from PIL import Image

I = np.asarray(PIL.Image.open('a.jpg'))
print(I)
im = PIL.Image.fromarray(np.uint8(I))
im.show()
