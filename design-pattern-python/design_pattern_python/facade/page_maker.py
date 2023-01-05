from asset import Database, HtmlWriter


class PageMaker:
    def __new__(cls):
        raise NotImplementedError("Cannot generate instance by constructor")

    @classmethod
    def make_welcome_page(cls, mailaddr: str, filename: str) -> None:
        mailprop = Database.get_properties("maildata")

        try:
            username = mailprop.get(mailaddr, None)
            f = open(filename, "w")
            writer = HtmlWriter(f)
            writer.title(f"{username}'s web page")
            writer.paragraph(f"Welcome to {username}'s web page!")
            writer.paragraph("Nice to meet you!")
            writer.mailto(mailaddr, username)
            writer.close()
            print(f"{filename} is created for {mailaddr} ({username})")

        except KeyError:
            print("specfied user does not exist!")
        except Exception:
            print("something wrong happend!")


if __name__ == "__main__":
    PageMaker.make_welcome_page("hyuki@example.com", "welcome.html")
