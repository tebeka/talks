import matplotlib.pyplot as plt
import numpy as np

xs = np.linspace(-1, 1)
c = plt.Circle((0, 0), 1, fill=False)
fig, ax = plt.subplots()
ax.set_aspect(1)
ax.add_artist(c)
ax.set_xlim(-1, 1)
ax.set_ylim(-1, 1)
ax.axhline()
ax.axvline()
ax.set_xticks([])
ax.set_yticks([])

n = 500
xs, ys = np.random.rand(n), np.random.rand(n)
mask, = np.where(np.sqrt(xs*xs+ys*ys) <= 1)
plt.scatter(xs[mask], ys[mask], color='g', marker='.')
mask, = np.where(np.sqrt(xs*xs+ys*ys) > 1)
plt.scatter(xs[mask], ys[mask], color='r', marker='.')
ax.text(0.35, -0.1, r'$r = 1$', fontsize=20)
plt.tight_layout()
fig.set_size_inches(10, 10)
# plt.show()
fig.savefig('sim_circle.png')
