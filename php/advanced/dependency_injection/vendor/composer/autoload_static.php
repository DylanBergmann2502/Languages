<?php

// autoload_static.php @generated by Composer

namespace Composer\Autoload;

class ComposerStaticInitdbe96fdd8fa2218d20081b5c66393872
{
    public static $prefixLengthsPsr4 = array (
        'P' => 
        array (
            'Psr\\Container\\' => 14,
        ),
    );

    public static $prefixDirsPsr4 = array (
        'Psr\\Container\\' => 
        array (
            0 => __DIR__ . '/..' . '/psr/container/src',
        ),
    );

    public static $classMap = array (
        'Composer\\InstalledVersions' => __DIR__ . '/..' . '/composer/InstalledVersions.php',
    );

    public static function getInitializer(ClassLoader $loader)
    {
        return \Closure::bind(function () use ($loader) {
            $loader->prefixLengthsPsr4 = ComposerStaticInitdbe96fdd8fa2218d20081b5c66393872::$prefixLengthsPsr4;
            $loader->prefixDirsPsr4 = ComposerStaticInitdbe96fdd8fa2218d20081b5c66393872::$prefixDirsPsr4;
            $loader->classMap = ComposerStaticInitdbe96fdd8fa2218d20081b5c66393872::$classMap;

        }, null, ClassLoader::class);
    }
}
