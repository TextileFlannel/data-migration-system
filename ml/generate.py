"""Генерация синтетических данных для обучения.

Идём «от эталона»: задаём эталонную реляционную схему, денормализуем её в
MongoDB-документы и запоминаем правильный ответ (метку) по каждому полю.

Выход (сырые данные, без признаков — признаки считает Go, `dms build-dataset`):
  docs.jsonl   — документы, по одному JSON на строку
  labels.json  — метки: путь поля -> решение (scalar_column / new_table_1toN / ...)
"""

from __future__ import annotations

import argparse


def main() -> None:
    parser = argparse.ArgumentParser(description="Генерация синтетики (документы + метки)")
    parser.add_argument("--out-docs", default="testdata/docs.jsonl")
    parser.add_argument("--out-labels", default="testdata/labels.json")
    parser.add_argument("--n", type=int, default=1000, help="сколько документов сгенерировать")
    args = parser.parse_args()

    # TODO: сгенерировать эталонные схемы, денормализовать в документы, собрать метки.
    raise NotImplementedError("генератор ещё не реализован")


if __name__ == "__main__":
    main()
