// Copyright 2022 Strims contributors
// SPDX-License-Identifier: AGPL-3.0-only

import React, { useEffect } from "react";
import sass, { Logger } from "sass";

import Nav from "../components/Nav";

const Foo: React.FC = () => {
  useEffect(() => {
    const scss = `
.chat {
  &__emote_container {
    // TODO: virus

    &--wide {
      transform: scaleX(1.5) scaleY(0.8);
      margin-left: calc(var(--width) * 0.25);
      margin-right: calc(var(--width) * 0.25);

      &.chat__emote_container--root {
        margin-left: calc(var(--width) * 0.25 + 3px);
        margin-right: calc(var(--width) * 0.25 + 3px);
      }
    }

    &--smol {
      transform: scaleX(0.5) scaleY(0.5);
      margin-left: calc(var(--width) * -0.25);
      margin-right: calc(var(--width) * -0.25);

      &.chat__emote_container--root {
        margin-left: calc(var(--width) * -0.25 + 3px);
        margin-right: calc(var(--width) * -0.25 + 3px);
      }
    }

    &--smol &--wide {
      margin: 0 calc(var(--width) * 0.125);
    }

    &--wide &--smol {
      margin: 0 calc(var(--width) * -0.375);
    }

    &--mirror {
      transform: scaleX(-1);
    }

    &--flip {
      transform: scaleY(-1);
    }

    &--spin {
      animation: spin linear 800ms 3;

      &:hover {
        animation: spin linear 800ms infinite;
      }
    }

    @keyframes spin {
      0% {
        transform: rotate(0deg);
      }
      100% {
        transform: rotate(360deg);
      }
    }

    @mixin rustle {
      animation: rustle linear 100ms 7;

      &.chat__emote_container--animated {
        animation: rustle linear 100ms
          calc(var(--animation-duration-ms) * var(--animation-iterations) / 100);
      }

      &:hover,
      &.chat__emote_container--animated:hover,
      &.chat__emote_container--animate_forever {
        animation: rustle linear 100ms infinite;
      }
    }

    @keyframes rustle {
      from {
        transform: translateX(0);
      }
      to {
        transform: translateX(3px);
      }
    }

    &--rustle {
      @include rustle;
    }

    &--gold {
      mask-image: var(--background-image);
      mask-size: var(--width) var(--height);
      mask-repeat: no-repeat;
      overflow: hidden;

      &.chat__emote_container--animated {
        mask-position: calc(-1 * var(--animation-end-on-frame) * var(--width)) 0;
        mask-size: var(--animation-spritesheet-width) var(--height);
      }

      &:before,
      &:after {
        top: 0;
        left: 0;
        display: block;
        position: absolute;
        content: "";
        height: 100%;
        width: 100%;
      }

      &:before {
        background: linear-gradient(
          45deg,
          hsl(44, 82%, 47%) 0%,
          hsl(49, 82%, 60%) 30%,
          hsl(58, 91%, 64%) 61%,
          hsl(58, 91%, 64%) 61%,
          hsl(57, 90%, 69%) 100%
        );
        filter: blur(2px);
      }

      &:after {
        animation: glimmer 5s linear 1;
        background: -webkit-radial-gradient(
          center,
          ellipse cover,
          rgba(255, 255, 255, 1) 1%,
          rgba(2, 2, 2, 1) 33%
        );
        mix-blend-mode: screen;
        opacity: 0;
        z-index: 2;
      }

      &:hover:after,
      &.chat__emote_container--animate_forever:after {
        animation: glimmer 5s linear infinite;
      }

      .chat__emote {
        filter: contrast(125%) grayscale(100%) brightness(0.75);
        margin: 0;
        mix-blend-mode: color-burn;
      }

      &.chat__emote_container--animated .chat__emote,
      &.chat__emote_container--animated .chat__emote--animate_forever {
        animation: none;
      }
    }

    @keyframes glimmer {
      0% {
        transform: translate(-50px, 50px) scale(4);
        opacity: 1;
      }
      20% {
        transform: translate(50px, -50px) scale(4);
        opacity: 1;
      }
      21% {
        opacity: 0;
      }
      100% {
        transform: translate(-50px, 50px) scale(4);
        opacity: 0;
      }
    }

    &--fast.chat__emote_container {
      .chat__emote,
      .chat__emote--animate_forever {
        animation-duration: calc(var(--animation-duration-ms) * 0.5ms);
      }

      .chat__emote {
        animation-iteration-count: calc(
          var(--animation-iterations) * 2 + var(--animation-end-on-frame) /
            var(--animation-frame-count)
        );
      }

      .chat__emote:hover,
      .chat__emote.chat__emote--animate_forever {
        animation-iteration-count: infinite;
      }
    }

    &--slow.chat__emote_container {
      .chat__emote,
      .chat__emote--animate_forever {
        animation-duration: calc(var(--animation-duration-ms) * 2ms);
      }

      .chat__emote {
        animation-iteration-count: calc(
          var(--animation-iterations) / 2 + var(--animation-end-on-frame) /
            var(--animation-frame-count)
        );
      }

      .chat__emote:hover,
      .chat__emote.chat__emote--animate_forever {
        animation-iteration-count: infinite;
      }
    }

    &--reverse .chat__emote.chat__emote--animated {
      animation-direction: var(--animation-reverse-direction);
    }

    &--pause .chat__emote.chat__emote--animated {
      animation: none;
    }

    &--lag {
      animation: lag 4s 2;

      &:hover {
        animation-iteration-count: infinite;
      }
    }

    @keyframes lag {
      0% {
        mask-image: linear-gradient(to top, transparent 100%, black 0%);
      }
      12% {
        mask-image: linear-gradient(to top, transparent 87.5%, black 0%);
      }
      25% {
        mask-image: linear-gradient(to top, transparent 75%, black 0%);
      }
      37% {
        mask-image: linear-gradient(to top, transparent 62.5%, black 0%);
      }
      50% {
        mask-image: linear-gradient(to top, transparent 50%, black 0%);
      }
      62% {
        mask-image: linear-gradient(to top, transparent 37.5%, black 0%);
      }
      75% {
        mask-image: linear-gradient(to top, transparent 25%, black 0%);
      }
      87% {
        mask-image: linear-gradient(to top, transparent 12.5%, black 0%);
      }
      100% {
        mask-image: linear-gradient(to top, transparent 0%, black 0%);
      }
    }

    @mixin floating-particles(
      $name,
      $particle-image: "",
      $particle-height: auto,
      $particle-width: auto,
      $particle-content: "",
      $particle-font-size: inherit,
      $animation-duration: 1.2s
    ) {
      &:before,
      &:after,
      & .chat__emote_container--#{$name}__extra_1:before,
      & .chat__emote_container--#{$name}__extra_1:after {
        @if $particle-image != "" {
          background-image: url($particle-image);
        }
        content: $particle-content;
        font-size: $particle-font-size;
        background-size: cover;
        display: block;
        height: $particle-height;
        width: $particle-width;
        position: absolute;
        bottom: 0;
        z-index: 3;
        visibility: hidden;
        animation: floating-particle $animation-duration 6;
      }

      &:before {
        /* left-outer particle */
        left: -8px;
        animation-delay: -($animation-duration - $animation-duration * 0.25);
      }

      &:after {
        /* right-inner particle */
        right: -2px;
      }

      & .chat__emote_container--#{$name}__extra_1 {
        &:before {
          /* left-inner particle */
          left: -3px;
          animation-delay: -($animation-duration - ($animation-duration * 0.75));
        }

        &:after {
          /* right-outer particle */
          right: -7px;
          animation-delay: -($animation-duration - $animation-duration * 0.5);
        }
      }

      &:hover,
      &.chat__emote_container--animate_forever {
        &:before,
        &:after,
        & .chat__emote_container--#{$name}__extra_1:before,
        & .chat__emote_container--#{$name}__extra_1:after {
          animation-iteration-count: infinite;
        }
      }
    }

    @keyframes floating-particle {
      0%,
      2% {
        opacity: 0;
      }
      0% {
        visibility: visible;
        transform: scale(0) translateY(calc(var(--height) * 0.5));
      }
      50% {
        opacity: 1;
      }
      100% {
        transform: scale(0.6) translateY(calc(var(--height) * -1.25));
        opacity: 0;
      }
    }

    &--jam {
      animation: jam 500ms 7;

      &:hover,
      &.chat__emote_container--animate_forever {
        animation-iteration-count: infinite;
      }

      @include floating-particles(
        $name: jam,
        $particle-image: "/assets/chat/modifiers/jam.svg",
        $particle-height: 22px,
        $particle-width: 22px
      );
    }

    @keyframes jam {
      0% {
        transform: translateY(0);
      }
      100% {
        transform: translateY(2px);
      }
    }

    &--love {
      &:hover,
      &.chat__emote_container--animate_forever {
        animation-iteration-count: infinite;
      }

      @include floating-particles($name: love, $particle-content: "💗", $particle-font-size: 18px);
    }

    &--banned {
      &:after {
        background-image: url("/assets/chat/modifiers/banned.png");
        content: "";
        display: block;
        height: 32px;
        left: calc(50% - 16px);
        pointer-events: none;
        position: absolute;
        top: 0px;
        width: 32px;
        z-index: 2;
      }
    }

    @mixin weather($name, $duration: 6s) {
      &:before,
      &:after {
        background-image: url("/assets/chat/modifiers/#{$name}-static.png");
        animation: #{$name} $duration 1;
      }

      @keyframes #{$name} {
        0% {
          background-image: url("/assets/chat/modifiers/#{$name}.png");
        }
        100% {
          background-image: url("/assets/chat/modifiers/#{$name}.png");
        }
      }
    }

    &--rain {
      @include weather("rain");
    }

    &--snow {
      @include weather("snow");
    }

    &--rain,
    &--snow {
      min-width: 48px;
      text-align: center;

      &:before,
      &:after {
        content: "";
        display: block;
        height: 100%;
        left: 0;
        pointer-events: none;
        position: absolute;
        top: 0;
        width: 100%;
        z-index: 2;
      }

      &:after {
        background-position: 32px 32px;
      }

      &:hover,
      &.chat__emote_container--animate_forever {
        &:before,
        &:after {
          animation-iteration-count: infinite;
        }
      }
    }

    &--worth {
      overflow: hidden;

      &:before,
      &:after {
        content: "💰";
        font-size: 16px;
        position: absolute;
        top: 0;
        animation: worth 0.6s 8;
        filter: drop-shadow(0px 0px 3px rgba(255, 255, 255, 0.7));
        transform: translateY(-25px);
        z-index: 3;
      }

      &:before {
        left: -3px;
      }

      &:after {
        right: -3px;
        animation-delay: 0.25s;
      }

      &:hover,
      &.chat__emote_container--animate_forever {
        &:before,
        &:after {
          animation-iteration-count: infinite;
        }
      }
    }

    @keyframes worth {
      0% {
        transform: translateY(-25px);
      }
      99% {
        transform: translateY(32px);
      }
      100% {
        transform: translateY(-25px);
      }
    }

    &--dank {
      min-width: 72px;
      text-align: center;

      &:before {
        background-image: url("/assets/chat/modifiers/dank.png");
        background-position: center;
        background-repeat: no-repeat;
        background-size: cover;
        animation: dank 1s 2;
        width: 100%;
        height: 100%;
        min-height: 32px;
        top: calc(max(32px - var(--height), 0px) * -0.5);
        left: 0;
        z-index: 0;
        content: "";
        position: absolute;
      }

      &:hover,
      &.chat__emote_container--animate_forever {
        &:before {
          animation-iteration-count: infinite;
        }
      }
    }

    @keyframes dank {
      0%,
      100% {
        filter: hue-rotate(0deg);
      }
      50% {
        filter: hue-rotate(360deg);
      }
    }

    &--pride {
      min-width: 48px;
      text-align: center;

      &:before {
        position: absolute;
        background: linear-gradient(
          to bottom,
          rgba(255, 0, 0, 0.8) 16%,
          rgba(255, 165, 0, 0.8) 16% 33%,
          rgba(255, 255, 0, 0.8) 33% 50%,
          rgba(0, 255, 0, 0.8) 50% 66%,
          rgba(0, 0, 255, 0.8) 66% 83%,
          rgba(75, 0, 130, 0.8) 83%
        );
        width: 100%;
        height: 100%;
        min-height: 32px;
        top: calc(max(32px - var(--height), 0px) * -0.5);
        left: 0;
        pointer-events: none;
        z-index: 0;
        content: "";
      }
    }

    &--hyper {
      filter: drop-shadow(0px 0px 3px rgba(255, 255, 255, 0.7));
      @include rustle;

      .chat__emote:after {
        display: block;
        content: "";
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        z-index: 2;

        background-image: inherit;
        background-size: inherit;
        background-position: inherit;
        background-repeat: no-repeat;

        filter: opacity(0.1) drop-shadow(0px 0px 0px red) drop-shadow(0px 0px 0px red)
          drop-shadow(0px 0px 0px red) drop-shadow(0px 0px 0px red);
        mix-blend-mode: color-dodge;
      }
    }

    &--angel {
      min-width: calc(max(110px, var(--width) + 60px));
      text-align: center;
      animation: angel-bounce 2s linear 3;

      &:before,
      &:after {
        content: "";
        background-image: url("/assets/chat/modifiers/angel.png");
        background-size: 65px 27px;
        height: 27px;
        width: 65px;
        position: absolute;
        top: -5px;
        transform-style: flat;
        z-index: -1;

        animation-iteration-count: 3.43;
        animation-timing-function: linear;
        animation-duration: 2s;
        transform-origin: right center;
      }

      &:before {
        animation-name: angel-wing-l;
        transform: scaleX(0.92) translateZ(-100px) rotateZ(0deg) rotateY(24deg) rotateX(0deg);
        left: -10px;
      }

      &:after {
        animation-name: angel-wing-r;
        transform: scaleX(-0.92) translateZ(-100px) rotateZ(0deg) rotateY(24deg) rotateX(0deg);
        right: 55px;
      }

      &:hover,
      &.chat__emote_container--animate_forever {
        &,
        &:before,
        &:after {
          animation-iteration-count: infinite;
        }
      }
    }

    &--smol &--angel {
      margin: 0 -20px;
    }

    &--wide &--angel {
      margin: 0 20px;
    }

    &--smol &--wide &--angel,
    &--wide &--smol &--angel {
      margin: 0 -10px;
    }

    @mixin angel-wing($side, $d) {
      @keyframes angel-wing-#{$side} {
        0% {
          transform: scaleX($d * 0.9) translateZ(-100px) rotateZ(40deg) rotateY(50deg)
            rotateX(60deg);
        }
        25% {
          transform: scaleX($d * 1.1) translateZ(-100px) rotateZ(25deg) rotateY(35deg)
            rotateX(50deg);
        }
        50% {
          transform: scaleX($d * 0.85) translateZ(-100px) rotateZ(-15deg) rotateY(20deg)
            rotateX(-20deg);
        }
        75% {
          transform: scaleX($d * 0.6) translateZ(-100px) rotateZ(10deg) rotateY(10deg)
            rotateX(35deg);
        }
        100% {
          transform: scaleX($d * 0.9) translateZ(-100px) rotateZ(40deg) rotateY(50deg)
            rotateX(65deg);
        }
      }
    }
    @include angel-wing("l", 1);
    @include angel-wing("r", -1);

    @keyframes angel-bounce {
      0% {
        transform: translateY(2px);
      }
      50% {
        transform: translateY(-2px);
      }
      100% {
        transform: translateY(2px);
      }
    }

    &--slide {
      clip-path: polygon(0 -100%, 100% -100%, 100% 200%, 0 200%);

      &__extra_1 {
        animation: slide 6s linear 1.5;

        @at-root &:hover,
          &.chat__emote_container--animate_forever {
          animation-iteration-count: infinite;
        }

        @at-root .chat__emote_container--reverse & {
          animation-direction: reverse;
        }

        @at-root .chat__emote_container--slow & {
          animation-duration: 12s;
        }

        @at-root .chat__emote_container--fase & {
          animation-duration: 3s;
        }
      }
    }

    @keyframes slide {
      0% {
        transform: translate(-110%, 0);
      }

      100% {
        transform: translate(110%, 0);
      }
    }

    &--peek {
      clip-path: polygon(0 -100%, 100% -100%, 100% 200%, 0 200%);

      &__extra_1 {
        animation: peek 6s ease-in-out 1.5;

        @at-root &:hover,
          &.chat__emote_container--animate_forever {
          animation-iteration-count: infinite;
        }

        @at-root .chat__emote_container--reverse & {
          animation-direction: reverse;
        }

        @at-root .chat__emote_container--slow & {
          animation-duration: 12s;
        }

        @at-root .chat__emote_container--fase & {
          animation-duration: 3s;
        }
      }
    }

    @keyframes peek {
      0% {
        transform: translate(-110%, 0);
      }

      50% {
        transform: translate(-0%, 0);
      }

      100% {
        transform: translate(-110%, 0);
      }
    }
  }
}
`;

    const compilationResult = sass.compileString(scss, {
      logger: Logger.silent,
      alertColor: false,
      sourceMap: false,
      quietDeps: true,
      functions: null,
      sourceMapIncludeSources: false,
      importers: null,
    });

    console.log(">>>", compilationResult);
  }, []);

  return <Nav />;
};

export default Foo;
