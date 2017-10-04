ConstantBase
===================

> This is experimental!

A constant time implementation of Base64, this is based on work of [Paragon Initiative Enterprises](https://github.com/paragonie/constant_time_encoding) and [Steve "Sc00bz" Thomas](https://github.com/Sc00bz/ConstTimeEncoding), both for PHP.

----------

Download
=======

    go get github.com/Inkeliz/ConstantBase/Base64

Usage
=======

For now we only support a Base64, to encode with padding use:

    Encoded := ConstantBase64.EncodeWithPad(Text)

If you don't care about padding use `Encode` instead:

    Encoded := ConstantBase64.Encode(Text)

In both cases, we use `Decode` to decode the Base64:

    Decoded := ConstantBase64.Decode(Encoded)

<<<<<<< HEAD
=======
    
>>>>>>> origin/master
