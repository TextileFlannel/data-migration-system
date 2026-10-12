"""Обучение модели решений.

Вход:  train.csv  — его собирает Go (`dms build-dataset`): признаки + метка.
Выход: models/model.txt — модель LightGBM в текстовом формате, которую Go читает
       через библиотеку leaves.
"""

from __future__ import annotations

import argparse


def main() -> None:
    parser = argparse.ArgumentParser(description="Обучение LightGBM по train.csv")
    parser.add_argument("--train", default="train.csv")
    parser.add_argument("--out", default="models/model.txt")
    args = parser.parse_args()

    # TODO: загрузить train.csv, обучить LightGBM, сохранить модель в текстовом формате.
    raise NotImplementedError("обучение ещё не реализовано")


if __name__ == "__main__":
    main()
