"""Оценка модели: метрики и сравнение с эвристическим baseline.

Считает accuracy / macro-F1 / confusion matrix и сравнивает обученную модель с
«тупым» эвристическим baseline на тех же данных.
"""

from __future__ import annotations

import argparse


def main() -> None:
    parser = argparse.ArgumentParser(description="Оценка модели и сравнение с baseline")
    parser.add_argument("--train", default="train.csv")
    parser.add_argument("--model", default="models/model.txt")
    args = parser.parse_args()

    # TODO: holdout/кросс-валидация, метрики, confusion matrix, сравнение с baseline.
    raise NotImplementedError("оценка ещё не реализована")


if __name__ == "__main__":
    main()
