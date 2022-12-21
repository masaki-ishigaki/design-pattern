class Singleton:
    __unique_instance = None

    def __new__(cls_):
        raise NotImplementedError("Cannot generate instance by constructor")

    @classmethod
    def __internal_new__(cls):
        return super().__new__(cls)

    @classmethod
    def get_instance(cls):
        if not cls.__unique_instance:
            cls.__unique_instance = cls.__internal_new__()
        return cls.__unique_instance


if __name__ == "__main__":
    print("Start")
    obj1 = Singleton.get_instance()
    obj2 = Singleton.get_instance()
    if obj1 == obj2:
        print("obj1とobj2は同じインスタンスです。")
    else:
        print("obj2とobj2は同じインスタンスではありません。")
    print("End")
