package main

import (
	"ImageProcessing/filter"
	"ImageProcessing/histogram"
	"ImageProcessing/io"
	"ImageProcessing/noise"
	"ImageProcessing/transform"
	"ImageProcessing/utils"
)

func main() {
	// open original.png
	original, err := io.Open("./img/original.png")
	if err != nil {
		panic(err)
	}

	// convert original.png to grayscale
	grayscale := filter.Grayscale(original)
	err = io.Save("./img/grayscale.png", grayscale, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of grayscale.png
	grayscalHistogram := histogram.GetGrayHistogram(grayscale)
	err = io.Save("./img/histogramImage/grayscalHistogram.png", grayscalHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Box Blur
	boxBlur := filter.BoxBlur(grayscale, 10)
	err = io.Save("./img/boxBlur.png", boxBlur, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of boxBlur.png
	boxBlurHistogram := histogram.GetGrayHistogram(boxBlur)
	err = io.Save("./img/histogramImage/boxBlurHistogram.png", boxBlurHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Gaussian Blur
	gaussianBlur := filter.GaussianBlur(grayscale, 3)
	err = io.Save("./img/gaussianBlur.png", boxBlur, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of boxBlur.png
	gaussianBlurHistogram := histogram.GetGrayHistogram(gaussianBlur)
	err = io.Save("./img/histogramImage/gaussianBlurHistogram.png", gaussianBlurHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Edge Detection with Laplacian
	edgeDetection := filter.EdgeDetection(grayscale, 5)
	err = io.Save("./img/edgeDetection.png", edgeDetection, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of edgeDetection.png
	edgeDetectionHistogram := histogram.GetGrayHistogram(edgeDetection)
	err = io.Save("./img/histogramImage/edgeDetectionHistogram.png", edgeDetectionHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Fixed Edge Detection
	fixedEdgeDetection := filter.FixedEdgeDetection(grayscale)
	err = io.Save("./img/fixEdgeDetection.png", fixedEdgeDetection, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of fixedEdgeDetection.png
	fixedEdgeDetectionHistogram := histogram.GetGrayHistogram(fixedEdgeDetection)
	err = io.Save("./img/histogramImage/fixedEdgeDetectionHistogram.png", fixedEdgeDetectionHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Sharpen
	sharpen := filter.Sharpen(grayscale, 2)
	err = io.Save("./img/sharpen.png", sharpen, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of sharpen.png
	sharpenHistogram := histogram.GetGrayHistogram(sharpen)
	err = io.Save("./img/histogramImage/sharpenHistogram.png", sharpenHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// FixedSharpen
	fixedSharpen := filter.FixedSharpen(grayscale)
	err = io.Save("./img/fixedSharpen.png", fixedSharpen, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of fixedSharpen.png
	fixedSharpenHistogram := histogram.GetGrayHistogram(fixedSharpen)
	err = io.Save("./img/histogramImage/fixedSharpenHistogram.png", fixedSharpenHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Median
	median := filter.Median(grayscale, 5)
	err = io.Save("./img/median.png", median, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of median.png
	medianHistogram := histogram.GetGrayHistogram(median)
	err = io.Save("./img/histogramImage/medianHistogram.png", medianHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Dilate
	dilate := filter.Dilate(grayscale, 5)
	err = io.Save("./img/dilate.png", dilate, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of dilate.png
	dilateHistogram := histogram.GetGrayHistogram(dilate)
	err = io.Save("./img/histogramImage/dilateHistogram.png", dilateHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Erode
	erode := filter.Erode(grayscale, 5)
	err = io.Save("./img/erode.png", erode, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of erode.png
	erodeHistogram := histogram.GetGrayHistogram(erode)
	err = io.Save("./img/histogramImage/erodeHistogram.png", erodeHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Uniform noise
	uniformNoise := noise.GenerateNoiseImage(original.Bounds().Dx(), original.Bounds().Dy(), noise.Uniform)
	err = io.Save("./img/uniformNoise.png", uniformNoise, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of uniformNoise.png
	uniformNoiseHistogram := histogram.GetGrayHistogram(uniformNoise)
	err = io.Save("./img/histogramImage/uniformNoiseHistogram.png", uniformNoiseHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Binary noise
	binaryNoise := noise.GenerateNoiseImage(original.Bounds().Dx(), original.Bounds().Dy(), noise.Binary)
	err = io.Save("./img/binaryNoise.png", binaryNoise, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of binaryNoise.png
	binaryNoiseHistogram := histogram.GetGrayHistogram(binaryNoise)
	err = io.Save("./img/histogramImage/binaryNoiseHistogram.png", binaryNoiseHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Gaussian noise
	gaussianNoise := noise.GenerateNoiseImage(original.Bounds().Dx(), original.Bounds().Dy(), noise.Gaussian)
	err = io.Save("./img/gaussianNoise.png", gaussianNoise, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of gaussianNoise.png
	gaussianNoiseHistogram := histogram.GetGrayHistogram(gaussianNoise)
	err = io.Save("./img/histogramImage/gaussianNoiseHistogram.png", gaussianNoiseHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Spike noise
	spikeNoise := noise.GenerateNoiseImage(original.Bounds().Dx(), original.Bounds().Dy(), noise.Spike)
	err = io.Save("./img/spikeNoise.png", spikeNoise, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of spikeNoise.png
	spikeNoiseHistogram := histogram.GetGrayHistogram(spikeNoise)
	err = io.Save("./img/histogramImage/spikeNoiseHistogram.png", spikeNoiseHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// grayscale.png with spike noise
	grayscaleWithSpikeNoise := noise.GenerateSpikeNoiseOn(grayscale)
	err = io.Save("./img/grayscaleWithSpikeNoise.png", grayscaleWithSpikeNoise, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of grayscaleWithSpikeNoise.png
	grayscaleWithSpikeNoiseHistogram := histogram.GetGrayHistogram(grayscaleWithSpikeNoise)
	err = io.Save("./img/histogramImage/grayscaleWithSpikeNoiseHistogram.png", grayscaleWithSpikeNoiseHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// smoothed grayscaleWithSpikeNoise.png by median filter
	smoothedGrayscaleWithSpikeNoiseByMedian := filter.Median(grayscaleWithSpikeNoise, 3)
	err = io.Save("./img/smoothedGrayscaleWithSpikeNoiseByMedian.png", smoothedGrayscaleWithSpikeNoiseByMedian, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of smoothedGrayscaleWithSpikeNoiseByMedian.png
	smoothedGrayscaleWithSpikeNoiseByMedianHistogram := histogram.GetGrayHistogram(smoothedGrayscaleWithSpikeNoiseByMedian)
	err = io.Save("./img/histogramImage/smoothedGrayscaleWithSpikeNoiseByMedianHistogram.png", smoothedGrayscaleWithSpikeNoiseByMedianHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Bilateral filter
	bilateralFilter := filter.BilateralFilter(grayscale, 5, 12, 16)
	err = io.Save("./img/bilateralFilter.png", bilateralFilter, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of smoothedGrayscaleWithSpikeNoiseByMedian.png
	bilateralFilterHistogram := histogram.GetGrayHistogram(bilateralFilter)
	err = io.Save("./img/histogramImage/bilateralFilterHistogram.png", bilateralFilterHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	bilateralFilter2 := filter.BilateralFilter(bilateralFilter, 5, 12, 16)
	err = io.Save("./img/bilateralFilter2.png", bilateralFilter2, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	bilateralFilter3 := filter.BilateralFilter(bilateralFilter2, 5, 12, 16)
	err = io.Save("./img/bilateralFilter3.png", bilateralFilter3, io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Gamma correct
	gammaCorrect := filter.Gamma(grayscale, 0.5)
	err = io.Save("./img/gammaCorrect.png", gammaCorrect, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of gammaCorrect.png
	gammaCorrectHistogram := histogram.GetGrayHistogram(gammaCorrect)
	err = io.Save("./img/histogramImage/gammaCorrectHistogram.png", gammaCorrectHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Histogram equalization
	histogramEqualization := filter.HistogramEqualization(grayscale)
	err = io.Save("./img/histogramEqualization.png", histogramEqualization, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of histogramEqualization.png
	histogramEqualizationHistogram := histogram.GetGrayHistogram(histogramEqualization)
	err = io.Save("./img/histogramImage/histogramEqualizationHistogram.png", histogramEqualizationHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// Cumulative histogram of histogramEqualization.png
	cumulativeHistogramEqualizationHistogram := histogram.GetGrayHistogram(histogramEqualization)
	err = io.Save("./img/histogramImage/cumulativeHistogramEqualizationHistogram.png", cumulativeHistogramEqualizationHistogram.Y.Cumulate().Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Halftoning with dithering method
	halftoningWithDitheringMethod := filter.HalftoningWithDitheringMethod(grayscale)
	err = io.Save("./img/halftoningWithDitheringMethod.png", halftoningWithDitheringMethod, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of halftoningWithDitheringMethod.png
	halftoningWithDitheringMethodHistogram := histogram.GetGrayHistogram(halftoningWithDitheringMethod)
	err = io.Save("./img/histogramImage/halftoningWithDitheringMethodHistogram.png", halftoningWithDitheringMethodHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// Halftoning with error diffusion method
	halftoningWithErrorDiffusionMethod := filter.HalftoningWithErrorDiffusionMethod(grayscale)
	err = io.Save("./img/halftoningWithErrorDiffusionMethod.png", halftoningWithErrorDiffusionMethod, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of halftoningWithDitheringMethod.png
	halftoningWithErrorDiffusionMethodHistogram := histogram.GetGrayHistogram(halftoningWithErrorDiffusionMethod)
	err = io.Save("./img/histogramImage/halftoningWithErrorDiffusionMethodHistogram.png", halftoningWithErrorDiffusionMethodHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// open keystoneEffectSample.png
	keystoneEffectSampleOriginal, err := io.Open("./img/keystoneEffectSample.png")
	if err != nil {
		panic(err)
	}

	// convert original.png to grayscale
	keystoneEffectSampleOriginalGrayscale := filter.Grayscale(keystoneEffectSampleOriginal)
	err = io.Save("./img/keystoneEffectSampleOriginalGrayscale.png", keystoneEffectSampleOriginalGrayscale, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of grayscale.png
	keystoneEffectSampleOriginalGrayscaleHistogram := histogram.GetGrayHistogram(keystoneEffectSampleOriginalGrayscale)
	err = io.Save("./img/histogramImage/keystoneEffectSampleOriginalGrayscaleHistogram.png", keystoneEffectSampleOriginalGrayscaleHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}

	// default
	ua, va := 790.0, 1138.0
	ub, vb := 1228.0, 1824.0
	uc, vc := 2626.0, 1190.0
	ud, vd := 2016.0, 744.0
	x, y := 400.0, 300.0
	a := [][]float64{
		// aの係数 bの係数 cの係数 dの係数 eの係数 fの係数  gの係数 		  hの係数 											   右辺
		{	 x, 	0, 		0, 	   0, 	  0, 	 0, 	 -ud*x, 		0,   		  										 ud-ua},
		{	 0, 	y, 		0, 	   0, 	  0, 	 0, 	 0, 			-ub*y,   	  										 ub-ua},
		{	 0, 	0, 		1, 	   0, 	  0, 	 0, 	 0, 			0,   		  										 ua},
		{	 0, 	0, 		0, 	   x, 	  0, 	 0, 	 -vd*x, 		0,   		  										 vd-va},
		{	 0, 	0, 		0, 	   0, 	  y, 	 0, 	 0, 			-vb*y,   	  										 vb-va},
		{	 0, 	0, 		0, 	   0, 	  0, 	 1, 	 0, 			0,   		  										 va},
		{	 0, 	0, 		0, 	   0, 	  0, 	 0, 	 -uc*x+ud*x, 	-uc*y+ub*y,   										 uc-ud-ub+ua},
		{	 0, 	0, 		0, 	   0, 	  0, 	 0, 	 0, 			-vc*y+vb*y-(-vc*x+vd*x)*(-uc*y+ub*y)/(-uc*x+ud*x),   vc-vd-vb+va-(-vc*x+vd*x)*(uc-ud-ub+ua)/(-uc*x+ud*x)},
	}

	utils.GaussElimination(&a)
	planeProjectionMatrix := [][]float64{
		{a[0][len(a[0])-1], a[1][len(a[0])-1], a[2][len(a[0])-1]},
		{a[3][len(a[0])-1], a[4][len(a[0])-1], a[5][len(a[0])-1]},
		{a[6][len(a[0])-1], a[7][len(a[0])-1], 1},
	}
	// fmt.Println(planeProjectionMatrix)
	// fmt.Println("A")
	// fmt.Println(planeProjectionMatrix[0][0]*0 + planeProjectionMatrix[0][1]*0 + planeProjectionMatrix[0][2]*1)
	// fmt.Println(planeProjectionMatrix[1][0]*0 + planeProjectionMatrix[1][1]*0 + planeProjectionMatrix[1][2]*1)
	// fmt.Println(planeProjectionMatrix[2][0]*0 + planeProjectionMatrix[2][1]*0 + planeProjectionMatrix[2][2]*1)
	// fmt.Println()
	// fmt.Println("B")
	// fmt.Println(planeProjectionMatrix[0][0]*0 + planeProjectionMatrix[0][1]*y + planeProjectionMatrix[0][2]*1)
	// fmt.Println(planeProjectionMatrix[1][0]*0 + planeProjectionMatrix[1][1]*y + planeProjectionMatrix[1][2]*1)
	// fmt.Println(planeProjectionMatrix[2][0]*0 + planeProjectionMatrix[2][1]*0 + planeProjectionMatrix[2][2]*1)
	// fmt.Println()
	// fmt.Println("C")
	// fmt.Println(planeProjectionMatrix[0][0]*x + planeProjectionMatrix[0][1]*y + planeProjectionMatrix[0][2]*1)
	// fmt.Println(planeProjectionMatrix[1][0]*x + planeProjectionMatrix[1][1]*y + planeProjectionMatrix[1][2]*1)
	// fmt.Println(planeProjectionMatrix[2][0]*0 + planeProjectionMatrix[2][1]*0 + planeProjectionMatrix[2][2]*1)
	// fmt.Println()
	// fmt.Println("D")
	// fmt.Println(planeProjectionMatrix[0][0]*x + planeProjectionMatrix[0][1]*0 + planeProjectionMatrix[0][2]*1)
	// fmt.Println(planeProjectionMatrix[1][0]*x + planeProjectionMatrix[1][1]*0 + planeProjectionMatrix[1][2]*1)
	// fmt.Println(planeProjectionMatrix[2][0]*0 + planeProjectionMatrix[2][1]*0 + planeProjectionMatrix[2][2]*1)

	// transform keystoneEffectSampleOriginalGrayscale.png
	keystoneEffect := transform.KeystoneEffect(keystoneEffectSampleOriginalGrayscale, &planeProjectionMatrix)
	err = io.Save("./img/keystoneEffect.png", keystoneEffect, io.PNGEncoder())
	if err != nil {
		panic(err)
	}
	// histogram of grayscale.png
	keystoneEffectHistogram := histogram.GetGrayHistogram(keystoneEffect)
	err = io.Save("./img/histogramImage/keystoneEffectHistogram.png", keystoneEffectHistogram.Y.Dump(), io.PNGEncoder())
	if err != nil {
		panic(err)
	}
}