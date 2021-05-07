# Distilling Your Examples

When we teach, we try to give realistic examples. The problem with realistic examples is that they tend to have code that is not related to the current teaching subject.

For example, when teaching about serialization with `pickle` - I started with the following code:

```python
@dataclass
class Location:
    lat: float
    lng: float

loc = Location(lat=48.8539241, lng=2.29133)
```

Which derailed the discussion to dataclasses instead of serialization.

In this talk, I'll present some code examples used in teaching. Then I'll show what are the parts that distracted the students and then show a distilled example which focuses on the subject matter.

---
- Simple IDE (popups ...)
