<?php
    class Matcher {
        private $image;
        private $st = "yes";
        public function setImage($url) {
            /* Attempt to open */
            $this->image = @imagecreatefrompng($url);
            imagealphablending($this->image, true);
            imagesavealpha($this->image, true);

            /* See if it failed */
            if(!$this->image)
            {
                echo "Image loading failed.";
            }
        }
        public function checkAtPoint($x, $y) {
            $rgb = imagecolorat($this->image, $x, $y);
            $r = ($rgb >> 16) & 0xFF;
            $g = ($rgb >> 8) & 0xFF;
            $b = $rgb & 0xFF;

            return sprintf("#%02x%02x%02x", $r, $g, $b); // #0d00ff
        }
        public function printImage() {
            // Set the content type header - in this case image/jpeg
            header('Content-type: image/png');

            // Output the image
            imagepng($this->image);

            // Free up memory
            imagedestroy($this->image);
        }
    }

    $matcher = new Matcher;
    $matcher->setImage("https://i.imgur.com/m4U4E7h.png");
    echo $matcher->checkAtPoint(8, 0);
    //$matcher->printImage();
?>