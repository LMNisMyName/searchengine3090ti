# coding:utf8
import os  # ,ipdb
import torch as t
import torchvision as tv
from model import CaptionModel
from config import Config
import requests as req
from PIL import Image
from io import BytesIO
import sys
import requests
IMAGENET_MEAN = [0.485, 0.456, 0.406]
IMAGENET_STD = [0.229, 0.224, 0.225]


def generate():
    opt = Config()
    #for k, v in kwargs.items():
    #    setattr(opt, k, v)
    device = t.device('cuda') if opt.use_gpu else t.device('cpu')

    # 数据预处理
    data = t.load(opt.caption_data_path, map_location=lambda s, l: s)
    word2ix, ix2word = data['word2ix'], data['ix2word']

    normalize = tv.transforms.Normalize(mean=IMAGENET_MEAN, std=IMAGENET_STD)
    transforms = tv.transforms.Compose([
        tv.transforms.Resize(opt.scale_size),
        tv.transforms.CenterCrop(opt.img_size),
        tv.transforms.ToTensor(),
        normalize
    ])

    response = req.get(sys.argv[1])
    
    img = Image.open(BytesIO(response.content))
    #img = Image.open(webimg)


    #img = Image.open(opt.test_img)

    img = transforms(img).unsqueeze(0)

    # 用resnet50来提取图片特征
    resnet50 = tv.models.resnet50(True).eval()
    del resnet50.fc
    resnet50.fc = lambda x: x
    resnet50.to('cpu')
    img = img.to('cpu')
    img_feats = resnet50(img).detach()

    # Caption模型
    model = CaptionModel(opt, word2ix, ix2word)
    model = model.load(opt.model_ckpt).eval()
    model.to('cpu')

    results = model.generate(img_feats.data[0])
    #for i in results:
    #    results.append(i.replace("</EOS>", ""))
    strlen=len(results[0])
    results[0]=results[0][:strlen-6]
    print(results[0])




if __name__ == '__main__':
    #import fire

    generate()
    #fire.Fire()
