import matplotlib.pyplot as plt
import numpy as np

def f1(x):
    return x*np.cos(x+5)

def f2(x):
    return 1/(1+25*x*x)

def read_data_from_file(filename):
    with open(filename, 'r') as file:
        n = int(file.readline().strip())
        nodes = [tuple(map(float, file.readline().strip().split())) for _ in range(n)]
        x_values = []
        y_values = []
        for _ in range(101):
            x, y = map(float, file.readline().strip().split())
            x_values.append(x)
            y_values.append(y)
    return nodes, x_values, y_values

def plot_graphs(nodes, x_values, y_values, filename):
    # Преобразование списка координат в отдельные списки x и y
    x_nodes, y_nodes = zip(*nodes)

    # Построение графика исходной функции
    x_original = np.arange(-5, 5, 0.1)
    y_original = f1(x_original)

    # Настройка графика
    plt.figure(figsize=(10, 6))
    plt.title('Graphs from File')
    plt.xlabel('X')
    plt.ylabel('Y')

    # Добавление узлов на график
    plt.scatter(x_nodes, y_nodes, color='blue', label='Nodes')

    # Добавление графика функции по точкам из файла
    plt.plot(x_values, y_values, color='red', label='Interpolated function')

    # Добавление графика исходной функции

    plt.plot(x_original, y_original, color='green', label='Original Function')

    plt.legend()
    plt.grid(True)
    plt.show()

# Имя файла с данными
filename = 'eq1.txt'

# Чтение данных из файла
nodes, x_values, y_values = read_data_from_file(filename)

# Построение графиков
plot_graphs(nodes, x_values, y_values, 'graphs_from_file.png')